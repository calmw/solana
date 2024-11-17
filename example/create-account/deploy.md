#### 编译

```shell
cargo-build-sbf --manifest-path ./program/Cargo.toml --sbf-out-dir ./program/target/so
```

#### 部署

```shell
solana program deploy ./program/target/so/program.so

# Program Id: KctJEm8gpQVDT1SmHSBW54WeSmrqmPEwRrS54eVzm9N
```