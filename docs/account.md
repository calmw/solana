#### account
- 

#### 创建账户

- ``` solana-keygen pubkey --outfile ./my-keypair2.json ``` 该方式只会生成一个账户地址并保存于文件中
- ``` solana-keygen new --outfile ./my-keypair.json  ``` 会生成一个账户并保存于文件中
    - 如果后面不加文件参数，会生成到默认的地址：
    - Byci3mh4DbgT1Dt7Ggp2Am16nT1NMkJwQn9HX2KNVfLp
    - BcdjTKkvd88Kkf79gKtHqMv43wgQHJvoV2CeeeMjRXYm
    - 一旦生成密钥对，您可以使用以下命令获取密钥对的地址（公钥）： ``` solana address ```
- 要验证您是否持有给定地址的私钥，请使用 solana-keygen verify
  ``` solana-keygen verify <PUBKEY> ~/my-solana-wallet/my-keypair.json ```
    - 例子 ``` solana-keygen verify 5N4iRJ5mncQEcPU1Cd1Nxm3hmVtDpQmZRtfoV5xgwq49  ./my-keypair.json ```

#### 水龙头领取测试币

- ``` solana airdrop 2 ```

#### 查看余额

- ``` solana balance SOME_WALLET_ADDRESS ```