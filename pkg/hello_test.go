package target
// package name should be same as hello.go


import (
   "mocktest/pkg/mocks"
   "testing"

   "github.com/golang/mock/gomock"
)

func TestHellobyChant(t *testing.T) {
    // create gomock controller
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    // Use mock interface 
    chant := himocks.NewMockmath(ctrl)

    // stub 
    chant.EXPECT().
          Add(gomock.Any(),gomock.Any()).
	  Return(2).
	  Times(1)

    // test
    v :=GaoAdd(chant, 2, 2)
    if v != 4 {
        t.Logf("Stub successful")
//	t.Fatal()
    }

}
