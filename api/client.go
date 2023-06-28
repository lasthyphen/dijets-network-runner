package api

import (
	"github.com/lasthyphen/utilitychain/plugin/evm"
	"github.com/lasthyphen/dijetsnode/api/admin"
	"github.com/lasthyphen/dijetsnode/api/health"
	"github.com/lasthyphen/dijetsnode/api/info"
	"github.com/lasthyphen/dijetsnode/api/ipcs"
	"github.com/lasthyphen/dijetsnode/api/keystore"
	"github.com/lasthyphen/dijetsnode/indexer"
	"github.com/lasthyphen/dijetsnode/vms/avm"
	"github.com/lasthyphen/dijetsnode/vms/platformvm"
)

// Issues API calls to a node
// TODO: byzantine api. check if appropriate. improve implementation.
type Client interface {
	PChainAPI() platformvm.Client
	XChainAPI() avm.Client
	XChainWalletAPI() avm.WalletClient
	CChainAPI() evm.Client
	CChainEthAPI() EthClient // ethclient websocket wrapper that adds mutexed calls, and lazy conn init (on first call)
	InfoAPI() info.Client
	HealthAPI() health.Client
	IpcsAPI() ipcs.Client
	KeystoreAPI() keystore.Client
	AdminAPI() admin.Client
	PChainIndexAPI() indexer.Client
	CChainIndexAPI() indexer.Client
	// TODO add methods
}
