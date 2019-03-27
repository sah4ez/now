## Contributors

Help wanted! 

For implementation clients in search for developers how have skills in Android, iOS, Web or Desktop development.

## Public storage public key

Specific storage for sharing public keys of pair ECDH by curve25519.

Service provide HTTP API for save and read a 32-bytes key with human-readable name (aka. alias).

This service solve problem for persistent saving public key and send to recipient through open communication channel.

### Features

- [x] save public key with name
- [x] load public key by name
- [ ] share some encrypted data via one-time link
- [ ] support other elliptic curve cryptography algorithms
- [ ] ephemeral encryption via API method
- [ ] verification of signature by link
- [ ] QR codes
- [ ] simple cli tool
- [ ] simple web-application
- [ ] simple iOS client
- [ ] simple android client
- [ ] local http server to pspk API

## Model

You should use this model:
```json
{
	"name":"Some Name",
	"key":"base64=="
}
```
where: 
- `name` - string with length more that 1000 signs
- `key` - encodeted to base64 32-bytes public key

## Request

Save public key:
```bash
curl -X POST "https://pspk.now.sh" -d '{"name":"Some.Name","key":"E7+TL112lj1GmJRHf9jT5MZJDgYIhUbtBLc4/ZFMZ5c="}'
{"access":true,"msg":"added"}
```

Read public key:
```bash
curl -X POST "https://pspk.now.sh" -d '{"name":"Some.Name"}'
{"access":true,"key":"wTaZA5+QeZpby33W2T5uV8TweWaPEZn3clTe5xkmb2M="}
```

## pspk cli usage

`pspk` - console tool which use API to pspk and implement encryption/decryption for one or several recipients.

```
$ pspk --help
NAME:
   pspk - pspk - encrypt you message and send through open communication channel

USAGE:
   pspk [global options] command [command options] [arguments...]

VERSION:
   0.1.2

DESCRIPTION:
   Console tool for encyption/decription data through pspk.now.sh

COMMANDS:
     publish, p                    --name <NAME> publish
     secret, s                     secret public_name
     encrypt, e                    ecnrypt pub_name some message will encrypt
     ephemeral-encrypt, ee         ee pub_name some message will encrypt
     decrypt, d                    decrypt pub_name base64==
     ephemeral-decrypt, ed         ephemeral-decryp pub_name base64==
     use-current, uc               --name name_pub_key use-current
     group, g                      --name base_name group
     start-group, sg               start-group groupName [pubName1 pubName2 ...]
     finish-group, fg              finish-group groupName pubName1 [pubName2 ...]
     secret-group, seg             secret-group groupName pubName1 [pubName2 ...]
     encrypt-group, eg             eg <GROUP_NAME> message
     ephemeral-encrypt-group, eeg  Encrypt input message with ephemeral key
     decrypt-group, dg             dg <GROUP_NAME> base64
     ephemeral-decrypt-group, edg  Decrypt input message with ephemral shared key
     help, h                       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --name value   key name
   --help, -h     show help
   --version, -v  print the version
```

### Generation key pair.
Will generation private and public keys and publish public pice to pspk.now.sh.
```bash
pspk --name <NAME_YOUR_KEY> publish
```

### Encryption some text message. 
Will encryption message through your private key and public key name from pspk.now.sh.
```bash
pspk --name <NAME_YOUR_KEY> encrypt <PUBLIC_PART> <SOME_MESSAGE_WITH_SPACES>
```

### Decription some text message. 
Will decription message through your private key and public key name from pspk.now.sh.
```bash
pspk --name <NAME_YOUR_KEY> decrypt <PUBLIC_PART> <SOME_BASE64_WITH_SPACES>
```

### Group encryption exchange
For encryption/decryption need generate shared secret in group. 
Use this algorithm (CLIQUES) [IV.A](https://pdfs.semanticscholar.org/dc45/970a9c43aaff17295c3769fdd0af9bded855.pdf) 

1. *Creat group*.
Create prime base point for group `base` and publish to pspk.now.sh
```
pspk --name base group
```
2. Decide number of members in group and select order for generation secret.
As example Alice, Bob, Carol and Daron want creage shared secret in group `base`.
```bash
pspk --name alice start-group base 
pspk --name bob start-group base alice
pspk --name carol start-group base bob alice
```
The last members finish generate intermediate secrets.
```bash
pspk --name daron finish-group base carol bob alice
```
Members can start generate shared secret keys via intermediate keys.
```bash
pspk --name daron secret-group base carol bob alice
pspk --name carol secret-group base daron bob alice
pspk --name bob secret-group base daron carol alice
pspk --name alice secret-group base daron carol bob
```
3. *Encryption* Encrypt some messages for `base` group members
```bash
pspk --name alice ephemeral-encrypt-group base Super secret message
```
4. *Decription* Decrypt the message from member's `base` group
```bash
pspk --name bob ephemeral-decryp-group base base64
```

**NOTE** All intermediate secrets would saved in pspk storage!

## pspk config

pspk use `$XDG_CONFIG_HOME` for saving configuration or default value `$HOME/.config/pspk`
Use `config.json` file for saving configuration:
```
{"current_name":"name"}
```

Also pspk use `$XDG_DATA_HOME` for saving appication data ro default value `$HOME/.local/share/pspk`
