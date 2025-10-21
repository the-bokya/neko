package application

import (
	"fmt"
	"neko/setup"
	"net/http"
	"os/exec"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type ResponseStatus string

const StatusOK ResponseStatus = "Ok"
const StatusNotOkay ResponseStatus = "Not Ok"

type Response struct {
	Status  ResponseStatus `json:"status"` // Either ok or not okay
	Message interface{}    `json:"message"`
	Data    interface{}    `json:"data"`
}
type VMSpec struct {
	Name     string `json:"name"`
	VCPUs    uint   `json:"vcpus"`
	Memory   uint   `json:"memory"`
	DiskSize uint   `json:"disk_size"`
	Image    string `json:"image"`
}

// passed in the response
type VMInfo struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

// TODO: Make everything that can be asynchronous asynchronous
// TODO: Cleanup
func (app *Application) CreateVM(c echo.Context) error {
	spec := &VMSpec{}
	if err := c.Bind(spec); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Status: StatusNotOkay, Message: err})
	}
	dom, err := app.Libvirt.Conn.LookupDomainByName(spec.Name)
	if err == nil {
		return c.JSON(http.StatusConflict, Response{Status: StatusNotOkay, Message: fmt.Sprintf("VM %s already exists. Please use another name", spec.Name)})
	}
	if spec.DiskSize < 5 {
		return c.JSON(http.StatusUnprocessableEntity, Response{Status: StatusNotOkay, Message: fmt.Sprintf("VM disk size of %dG is too small. It should be at least 5G", spec.DiskSize)})
	}
	vmImage := app.getVMImageFromName(spec.Image)
	if vmImage == nil {
		return c.JSON(http.StatusNotFound, Response{Status: StatusNotOkay, Message: fmt.Sprintf("Image %s not found", spec.Image)})
	}
	baseImagePath := vmImage.GetBaseImagePath()
	targetPath := filepath.Join(setup.EtcPath, "disks", fmt.Sprintf("%s.qcow2", spec.Name))
	if err := exec.Command("qemu-img", "convert", "-O", "qcow2", baseImagePath, targetPath).Run(); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: StatusNotOkay, Message: err})
	}
	if err := exec.Command("qemu-img", "resize", targetPath, fmt.Sprintf("%dG", spec.DiskSize)).Run(); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: StatusNotOkay, Message: err})
	}
	dom, err = app.Libvirt.CreateVM(spec.Name, targetPath, spec.VCPUs, spec.Memory)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: StatusNotOkay, Message: err})
	}
	dom.Free()
	dom, err = app.Libvirt.Conn.LookupDomainByName(spec.Name)
	if err != nil {
		fmt.Println(err)
	}
	domainUUID, _ := dom.GetUUIDString()
	domainName, _ := dom.GetName()
	dom.Free()
	return c.JSON(http.StatusOK, Response{Status: StatusOK, Message: "VM created successfully", Data: VMInfo{Name: domainName, UUID: domainUUID}})
}

func (app *Application) getVMImageFromName(name string) *setup.VMImage {
	for _, vmImage := range app.VMConfig.VMImages {
		if vmImage.Name == name {
			return &vmImage
		}
	}
	return nil
}
