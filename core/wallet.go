package core

import (
	"taskpot.com/go_wallet/model"
)

func getWalletById(id int64) (*model.Wallet, error) {
	//helper.Logger().Info("Fetching wallet from id" + strconv.FormatInt(id, 10))
	//res := database.GetWalletById(id)
	//var wallet *model.Wallet
	//err := res.Scan(&wallet)
	//return wallet, err
	return nil, nil
}

func CreateWallet(wallet *model.Wallet) (*model.Wallet, error){
	//helper.Logger().Info("Creating wallet for " + wallet.EntityUuid + " of type " + wallet.EntityType)
	//helper.Logger().Info("Wallet type " + wallet.WalletType)
	//res,err := database.CreateWallet(wallet)
	//if err != nil {
	//	return nil, err
	//}
	//id, err := res.LastInsertId()
	//if err != nil {
	//	return nil, err
	//}
	//return getWalletById(id)
	return nil, nil
}
