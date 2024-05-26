package registery

import (
	c "htmx-engineg/endpoints/controllers"
	hello_controller "htmx-engineg/endpoints/controllers/hello"
)

var ControllerRegistery = []c.Controller{
	hello_controller.HelloController,
}
