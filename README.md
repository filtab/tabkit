# tabkit
filtabOS operation toolkit go-sdk

# quick peek
When using mine toolkit you're connecting API service from filtabOS cloud.
All params and action informations are protected by EDCURVE-25519 signature and HMAC-SHA256 plein text encryption along with TLS.


# add dependency underlying to go.mod
github.com/filtab/tabkit v0.0.1

#### mining tools
### you'll have ability to withdraw your FIL/BTC/ETH balance e.g to your 3rd party wallet or bare address by your filtabOS account(or token instead)


    // filtabOS official usages
    import "github.com/filtab/tabkit/mine"
    
    // create your client
    cli := mine.CreateByToken("SFInnoscnuHUISfphXXjoiojisg")
    
    // get information of your account
    info := cli.Info()
    
    // get balance of your miners
    balance := cli.GetBalance(mine.FIL)
    
    // withdraw some of them 
    ack, err := cli.WithdrawTx(mine.FIL, 100, "t3vrkrtcyalf3766qzcudytrud7jvxzi37dskib3thkgmdmxzq34moglvk7rfpypxcpuosy5hxjbeanlf654sq")
   