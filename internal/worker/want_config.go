package worker

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	pb "github.com/loranode/gateway/api/meshtastic"
)

// wantConfig builds the ToRadio frame that asks the node to dump its config and
// full node database.
func (*Worker) wantConfig() ([]byte, error) {
	b, err := proto.Marshal(&pb.ToRadio{
		PayloadVariant: &pb.ToRadio_WantConfigId{WantConfigId: wantConfigID},
	})
	if err != nil {
		return nil, fmt.Errorf("marshal want_config: %w", err)
	}

	return b, nil
}
