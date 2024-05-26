package registery

import (
	c "htmx-engineg/src/common/controller"
	editController "htmx-engineg/src/controllers/edit"
	helloController "htmx-engineg/src/controllers/hello"
)

var ControllerRegistery = []c.Controller{
	helloController.Controller,
	editController.Controller,
}
