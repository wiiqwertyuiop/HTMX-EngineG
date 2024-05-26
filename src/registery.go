package registery

import (
	editController "htmx-engineg/src/controllers/edit"
	helloController "htmx-engineg/src/controllers/hello"
	c "htmx-engineg/src/models/controller"
)

var ControllerRegistery = []c.Controller{
	helloController.Controller,
	editController.Controller,
}
