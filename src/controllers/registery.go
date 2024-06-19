package regiserty

import (
	c "htmx-engineg/src/common/controllerRouter"
	editController "htmx-engineg/src/controllers/contact"
	helloController "htmx-engineg/src/controllers/hello"
)

var ControllerRegistery = []c.Controller{
	helloController.Controller,
	editController.Controller,
}
