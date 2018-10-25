package main

import (
"testing"
)

func TestCheckforsubset(t *testing.T) {

  tables := map[string]struct {
    		s1 string
		s2 string
                expectedOutput bool
	    }{
	    "test1": {
                 s1: "abcd",
		 s2: "abc",
                 expectedOutput: true,
		},
            "test2":{
                 s1: "abc",
                 s2: "abcd",
                 expectedOutput: false,
	},
     }
  for desc, ut := range tables {
   obtainedOutput:= checkForSubset(ut.s1, ut.s2)
   if ut.expectedOutput != obtainedOutput {
       t.Errorf("Test case failure: %v", desc)
   }
}
}

func TestRemoveDuplicates(t *testing.T) {
  tables := map[string]struct {
		elements []string
		expectedOutput []string
		}{
		"test1": {
			elements: []string{"abc","def","def"},
			expectedOutput: []string{"abc","def"},
		},
	}
 for desc, ut := range tables {
  obtainedOutput := removeDuplicates(ut.elements)
  if ut.expectedOutput != obtainedOutput {
	t.Errorf("Test case failure: %v", desc)
    }
}
}


