package live_migration_controller

import (
	"context"
	"fmt"
	"github.com/checkpoint-restore/go-criu/v6"
	"github.com/checkpoint-restore/go-criu/v6/rpc"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"log"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strconv"
)

// Reconciler reconciles a ShadowPod object.
type Reconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// Reconcile  objects.
// func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
// }

func testCRIUenv(ctx context.Context) {

	// CRIU local test
	c := criu.MakeCriu()
	version, err := c.GetCriuVersion()
	if err != nil {
		klog.ErrorS(err, "CRIU not installed, that's a bummer")
	} else {
		klog.Infof("hooray! CRIU installed, version: %s", version)
	}

	// make the actual dump
	err = doDump(c, "1", "/tmp/criu", false, "")
	if err != nil {
		klog.ErrorS(err, "doDump error")

	}

	/* client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		klog.ErrorS(err, "containerd not installed, that's a bummer")
	}
	defer client.Close()

	// Checkpoint a container
	ctx = context.Background()
	container, err := client.LoadContainer(ctx, "1")
	if err != nil {
		klog.ErrorS(err, "container not found")
	}
	checkpoint, err := container.Checkpoint(ctx, "checkpoint_name")
	if err != nil {
		klog.ErrorS(err, "Error in Checkpoint")
	}

	size, err := checkpoint.Size(ctx)
	if err != nil {
		klog.ErrorS(err, "Error in Size, probably checkpoint does not exists")
	}
	klog.Infof("checkpoint size: %d", size) */
}

func doDump(c *criu.Criu, pidS string, imgDir string, pre bool, prevImg string) error {
	klog.Infof("Dumping")
	pid, err := strconv.ParseInt(pidS, 10, 32)
	if err != nil {
		return fmt.Errorf("can't parse pid: %w", err)
	}
	img, err := os.Open(imgDir)
	if err != nil {
		return fmt.Errorf("can't open image dir: %w", err)
	}

	defer func(img *os.File) {
		err := img.Close()
		if err != nil {

		}
	}(img)

	opts := &rpc.CriuOpts{
		Pid:         proto.Int32(int32(pid)),
		ImagesDirFd: proto.Int32(int32(img.Fd())),
		LogLevel:    proto.Int32(4),
		LogFile:     proto.String("dump.log"),
	}

	if prevImg != "" {
		opts.ParentImg = proto.String(prevImg)
		opts.TrackMem = proto.Bool(true)
	}

	if pre {
		err = c.PreDump(opts, TestNfy{})
	} else {
		err = c.Dump(opts, TestNfy{})
	}
	if err != nil {
		return fmt.Errorf("dump fail: %w", err)
	}

	return nil
}

// TestNfy struct
type TestNfy struct {
	criu.NoNotify
}

// PreDump test function
func (c TestNfy) PreDump() error {
	log.Println("TEST PRE DUMP")
	return nil
}
