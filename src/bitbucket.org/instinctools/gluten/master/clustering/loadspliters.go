package projectmodifiers

import (
	core "bitbucket.org/instinctools/gluten/core"
)

type LoadSplitter interface {
	
	Split(step core.TestStep) []core.TestStep
	
}

