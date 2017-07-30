# SafeWord

## SafeWord Website

https://safeword.airshipcms.io/

Source : https://github.com/naotoyamaguchi/safeword-website


## Scripts

`cd` into `./scripts`

note: ips are hardcoded atm

```sh
./snapshot.sh
```
creates a `scripts/cam/` directory that saves snapshots from pi camera
```sh
./record.sh
```
creates a `scripts/videos/` directory that saves videos from pi camera

recording has some settings, for record length.


```sh
./motionDetection.sh
```
uploads all motion detected snapshots from `~/motion` to s3

recording has some settings, for record length.
