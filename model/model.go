package model

type TickerALL struct {
	Tickers struct {
		BtcIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolBtc     string `json:"vol_btc"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"btc_idr"`
		TenIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolTen     string `json:"vol_ten"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"ten_idr"`
		OneInchIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			Vol1Inch   string `json:"vol_1inch"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"1inch_idr"`
		AaveIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolAave    string `json:"vol_aave"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"aave_idr"`
		AbyssIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolAbyss   string `json:"vol_abyss"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"abyss_idr"`
		ActIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolAct     string `json:"vol_act"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"act_idr"`
		AdaIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolAda     string `json:"vol_ada"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"ada_idr"`
		AlgoIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolAlgo    string `json:"vol_algo"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"algo_idr"`
		AoaIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolAoa     string `json:"vol_aoa"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"aoa_idr"`
		AtomIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolAtom    string `json:"vol_atom"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"atom_idr"`
		AttIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolAtt     string `json:"vol_att"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"att_idr"`
		BalIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolBal     string `json:"vol_bal"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"bal_idr"`
		BatIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolBat     string `json:"vol_bat"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"bat_idr"`
		BcdIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolBcd     string `json:"vol_bcd"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"bcd_idr"`
		BchIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolBch     string `json:"vol_bch"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"bch_idr"`
		BchaIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolBcha    string `json:"vol_bcha"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"bcha_idr"`
		BnbIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolBnb     string `json:"vol_bnb"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"bnb_idr"`
		BsvIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolBsv     string `json:"vol_bsv"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"bsv_idr"`
		BtgIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolBtg     string `json:"vol_btg"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"btg_idr"`
		BtsIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolBts     string `json:"vol_bts"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"bts_idr"`
		BttIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolBtt     string `json:"vol_btt"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"btt_idr"`
		CelIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolCel     string `json:"vol_cel"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"cel_idr"`
		CeloIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolCelo    string `json:"vol_celo"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"celo_idr"`
		CkbIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolCkb     string `json:"vol_ckb"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"ckb_idr"`
		CoalIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolCoal    string `json:"vol_coal"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"coal_idr"`
		CompIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolComp    string `json:"vol_comp"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"comp_idr"`
		CotiIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolCoti    string `json:"vol_coti"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"coti_idr"`
		CroIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolCro     string `json:"vol_cro"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"cro_idr"`
		CrvIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolCrv     string `json:"vol_crv"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"crv_idr"`
		DadIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolDad     string `json:"vol_dad"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"dad_idr"`
		DaiIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolDai     string `json:"vol_dai"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"dai_idr"`
		DashIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolDash    string `json:"vol_dash"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"dash_idr"`
		DaxIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolDax     string `json:"vol_dax"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"dax_idr"`
		DepIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolDep     string `json:"vol_dep"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"dep_idr"`
		DgbIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolDgb     string `json:"vol_dgb"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"dgb_idr"`
		DgxIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolDgx     string `json:"vol_dgx"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"dgx_idr"`
		DogeIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolDoge    string `json:"vol_doge"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"doge_idr"`
		DotIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolDot     string `json:"vol_dot"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"dot_idr"`
		EgldIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolEgld    string `json:"vol_egld"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"egld_idr"`
		EmIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolEm      string `json:"vol_em"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"em_idr"`
		EnjIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolEnj     string `json:"vol_enj"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"enj_idr"`
		EosIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolEos     string `json:"vol_eos"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"eos_idr"`
		EtcIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolEtc     string `json:"vol_etc"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"etc_idr"`
		EthIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolEth     string `json:"vol_eth"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"eth_idr"`
		EursIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolEurs    string `json:"vol_eurs"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"eurs_idr"`
		FilIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolFil     string `json:"vol_fil"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"fil_idr"`
		FiroIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolFiro    string `json:"vol_firo"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"firo_idr"`
		GrtIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolGrt     string `json:"vol_grt"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"grt_idr"`
		GscIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolGsc     string `json:"vol_gsc"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"gsc_idr"`
		GxcIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolGxc     string `json:"vol_gxc"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"gxc_idr"`
		HbarIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolHbar    string `json:"vol_hbar"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"hbar_idr"`
		HedgIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolHedg    string `json:"vol_hedg"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"hedg_idr"`
		HiveIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolHive    string `json:"vol_hive"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"hive_idr"`
		HnstIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolHnst    string `json:"vol_hnst"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"hnst_idr"`
		HotIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolHot     string `json:"vol_hot"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"hot_idr"`
		HpbIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolHpb     string `json:"vol_hpb"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"hpb_idr"`
		IdkIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolIdk     string `json:"vol_idk"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"idk_idr"`
		IgnisIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolIgnis   string `json:"vol_ignis"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"ignis_idr"`
		IostIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolIost    string `json:"vol_iost"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"iost_idr"`
		IotaIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolIota    string `json:"vol_iota"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"iota_idr"`
		JstIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolJst     string `json:"vol_jst"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"jst_idr"`
		KdagIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolKdag    string `json:"vol_kdag"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"kdag_idr"`
		LetIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolLet     string `json:"vol_let"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"let_idr"`
		LgoldIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolLgold   string `json:"vol_lgold"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"lgold_idr"`
		LinkIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolLink    string `json:"vol_link"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"link_idr"`
		LlandIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolLland   string `json:"vol_lland"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"lland_idr"`
		LsilverIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolLsilver string `json:"vol_lsilver"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"lsilver_idr"`
		LtcIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolLtc     string `json:"vol_ltc"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"ltc_idr"`
		LyfeIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolLyfe    string `json:"vol_lyfe"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"lyfe_idr"`
		ManaIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolMana    string `json:"vol_mana"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"mana_idr"`
		MaticIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolMatic   string `json:"vol_matic"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"matic_idr"`
		MblIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolMbl     string `json:"vol_mbl"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"mbl_idr"`
		MkrIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolMkr     string `json:"vol_mkr"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"mkr_idr"`
		NeoIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolNeo     string `json:"vol_neo"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"neo_idr"`
		NrgIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolNrg     string `json:"vol_nrg"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"nrg_idr"`
		NxtIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolNxt     string `json:"vol_nxt"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"nxt_idr"`
		OceanIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolOcean   string `json:"vol_ocean"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"ocean_idr"`
		OctoIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolOcto    string `json:"vol_octo"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"octo_idr"`
		OkbIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolOkb     string `json:"vol_okb"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"okb_idr"`
		OmgIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolOmg     string `json:"vol_omg"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"omg_idr"`
		OntIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolOnt     string `json:"vol_ont"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"ont_idr"`
		OrbsIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolOrbs    string `json:"vol_orbs"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"orbs_idr"`
		OxtIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolOxt     string `json:"vol_oxt"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"oxt_idr"`
		PaxgIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolPaxg    string `json:"vol_paxg"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"paxg_idr"`
		QtumIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolQtum    string `json:"vol_qtum"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"qtum_idr"`
		RenIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolRen     string `json:"vol_ren"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"ren_idr"`
		RepIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolRep     string `json:"vol_rep"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"rep_idr"`
		RvnIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolRvn     string `json:"vol_rvn"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"rvn_idr"`
		SandIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolSand    string `json:"vol_sand"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"sand_idr"`
		SfiIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolSfi     string `json:"vol_sfi"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"sfi_idr"`
		SnxIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolSnx     string `json:"vol_snx"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"snx_idr"`
		SolveIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolSolve   string `json:"vol_solve"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"solve_idr"`
		SumoIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolSumo    string `json:"vol_sumo"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"sumo_idr"`
		SushiIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolSushi   string `json:"vol_sushi"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"sushi_idr"`
		TadIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolTad     string `json:"vol_tad"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"tad_idr"`
		TfuelIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolTfuel   string `json:"vol_tfuel"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"tfuel_idr"`
		ThetaIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolTheta   string `json:"vol_theta"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"theta_idr"`
		TitanIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolTitan   string `json:"vol_titan"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"titan_idr"`
		TrxIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolTrx     string `json:"vol_trx"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"trx_idr"`
		UmaIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolUma     string `json:"vol_uma"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"uma_idr"`
		UniIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolUni     string `json:"vol_uni"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"uni_idr"`
		UsdcIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolUsdc    string `json:"vol_usdc"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"usdc_idr"`
		UsdtIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolUsdt    string `json:"vol_usdt"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"usdt_idr"`
		VexIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolVex     string `json:"vol_vex"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"vex_idr"`
		VidyIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolVidy    string `json:"vol_vidy"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"vidy_idr"`
		VidyxIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolVidyx   string `json:"vol_vidyx"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"vidyx_idr"`
		VraIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolVra     string `json:"vol_vra"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"vra_idr"`
		VsysIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolVsys    string `json:"vol_vsys"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"vsys_idr"`
		WavesIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolWaves   string `json:"vol_waves"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"waves_idr"`
		WbtcIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolWbtc    string `json:"vol_wbtc"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"wbtc_idr"`
		WnxmIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolWnxm    string `json:"vol_wnxm"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"wnxm_idr"`
		WozxIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolWozx    string `json:"vol_wozx"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"wozx_idr"`
		XdcIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolXdc     string `json:"vol_xdc"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"xdc_idr"`
		XemIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolXem     string `json:"vol_xem"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"xem_idr"`
		XlmIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolXlm     string `json:"vol_xlm"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"xlm_idr"`
		XmrIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolXmr     string `json:"vol_xmr"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"xmr_idr"`
		XrpIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolXrp     string `json:"vol_xrp"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"xrp_idr"`
		XsgdIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolXsgd    string `json:"vol_xsgd"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"xsgd_idr"`
		XtzIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolXtz     string `json:"vol_xtz"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"xtz_idr"`
		YfiIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolYfi     string `json:"vol_yfi"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"yfi_idr"`
		YfiiIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolYfii    string `json:"vol_yfii"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"yfii_idr"`
		ZecIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolZec     string `json:"vol_zec"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"zec_idr"`
		ZilIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolZil     string `json:"vol_zil"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"zil_idr"`
		ZrxIdr struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolZrx     string `json:"vol_zrx"`
			VolIdr     string `json:"vol_idr"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"zrx_idr"`
		BtcUsdt struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolBtc     string `json:"vol_btc"`
			VolUsdt    string `json:"vol_usdt"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"btc_usdt"`
		BttUsdt struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolBtt     string `json:"vol_btt"`
			VolUsdt    string `json:"vol_usdt"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"btt_usdt"`
		EthUsdt struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolEth     string `json:"vol_eth"`
			VolUsdt    string `json:"vol_usdt"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"eth_usdt"`
		GardUsdt struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolGard    string `json:"vol_gard"`
			VolUsdt    string `json:"vol_usdt"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"gard_usdt"`
		PundixUsdt struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolPundix  string `json:"vol_pundix"`
			VolUsdt    string `json:"vol_usdt"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"pundix_usdt"`
		PxgUsdt struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolPxg     string `json:"vol_pxg"`
			VolUsdt    string `json:"vol_usdt"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"pxg_usdt"`
		SspUsdt struct {
			High       string `json:"high"`
			Low        string `json:"low"`
			VolSsp     string `json:"vol_ssp"`
			VolUsdt    string `json:"vol_usdt"`
			Last       string `json:"last"`
			Buy        string `json:"buy"`
			Sell       string `json:"sell"`
			ServerTime int    `json:"server_time"`
		} `json:"ssp_usdt"`
	} `json:"tickers"`
}

type DetailsTicker struct {
	Ticker struct {
		High       string `json:"high"`
		Low        string `json:"low"`
		VolTen     string `json:"vol_ten"`
		VolIdr     string `json:"vol_idr"`
		Last       string `json:"last"`
		Buy        string `json:"buy"`
		Sell       string `json:"sell"`
		ServerTime int    `json:"server_time"`
	} `json:"ticker"`
}
