package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const fireblockPubKey = `
-----BEGIN PGP PUBLIC KEY BLOCK-----

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
0cppZHefPO8ogUv/Dt4ayVQ8ug==
=TE5J
-----END PGP PUBLIC KEY BLOCK-----
`

const fireblockPrivKey = `
-----BEGIN PGP PRIVATE KEY BLOCK-----

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
=CVvx
-----END PGP PRIVATE KEY BLOCK-----
`

const signatureDetached = `-----BEGIN PGP SIGNATURE-----
Version: OpenPGP.js v2.5.5
Comment: https://openpgpjs.org

wsFcBAEBCAAQBQJZeKN9CRDNWDSj4H7bLAAAMN0P/iXRrPWP1zSFC08Q4fLO
tFuW+2PjOAk4Aa0IIHf9oVAEXlDZsO3fpayrv6Kuxsk18WQBn1zgzrTNPf0v
LjKZLRPak7tUyOcWB6VXHJ/hu4K2gl4AKkraiYew8tsjHuLeMJEpkZJCXZhI
lVsiONZJNLayEWcbAtp/zsmY96VTybu2/I53uclQXMRlHJuZAKgHO311grQ5
1sood+GIxhevnPM6Lll1P3AjHJulfp1rPCvg386dTJbjSHx+WlBxLgUpcqpg
S3YAWORqrfsEVCCKUzkC2aw3ld8FUAGqeDjMdy9uxuAQqqzWSwcg7ukpi9RJ
sULYzDysvETs2xxuylAH3mVSXclv3kgacoBNKwrF6i/jRZVcVGL7tj0sixDs
nFU8A25dzP7BmkD4N8e11czpvrDOqKFzhw5ZNW5dAflEoJJW78MSR3BWV4fu
T1E2yJH7LuLSOt1eIKr8+/xdkRf9dOJ1iPFwu6TdCsgzHmwW5PvqG11J770q
fW/2YDGOtaEMba40UklgTwl9dvBgG15hdQsjH2La/vGVb3/oiZSsLfut13ir
/jOno+F8FMdyrfjkObOWKV4bxxZEq6PTWmBK6NXMbTG34qKFIA+52NIVZiOW
w+GEALRW6XUiYy3OhV1/f80T3E+ALkYiCNhPxoBB5oUWKuC+2GcMDVnfK/1Z
oTE/
=NQOe
-----END PGP SIGNATURE-----`

const signatureClear = `-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA256

message from fireblock
-----BEGIN PGP SIGNATURE-----

iQIcBAEBCAAGBQJZ0RxKAAoJEM1YNKPgftsss+YP/1JiZmmZSMWAEziWnwJoM9Ec
81wamgqP2E4fxCSpD+Cnp51PKwr+vkz9UxR9eKHVzsO1RM+jr5mgFtXQa8Uy7tYf
LwVApbFDNIqQo4wvQq/lx7ydnRxW6A59drkaObVVthb8torECNfmgD5MuUx3KtIK
4pZct/9NSJKQ6V0PF9WX3l+uVYzBkEVmgm3oGR5UKesDHrE2dzspubD4vITzQq44
Bm+BSZhTNwSE1HwhwiwAddaa0MZ/48dEZsR3YQ2JbtLXQt8mmD89VcH0dE5BOPoU
L/42eAGgkSiPARncSGo+4iRxVsYhRYnKRnTGoac9EoishdozZeBb8kepCd0BW9K2
Q0aojj9NlfQ2bLYPFJVYKI67hjdyf05+oXPA9h2U7fkt5x2fbmc/0WSStheLKjnJ
2faWTxIJ7APPKYGQOTWZfRZJygIHcSgNUuXC3Rs/1YO5qEHJkWpLmZ1w5k27W2Ah
kqX4SGZufPDP/YNSUPFRkmlM6IdLFTkGUhXzqQeKq2mwL5+v0ptfttQr00ZU0qOd
xzPcZCSsgsQc+YbQxLD0M8Qj6k8iB2t2z0KJF0m7yfdo35UC7lyzvYX9tPd2l7Ag
0o1JRYa5BsHDA11qlH106Yl6h1hQwsczpXkE8xxI6Ij3QmUmhHsK+viWgGzwEiRp
S/LhAJEZEsxivJBlbeA1
=XhnQ
-----END PGP SIGNATURE-----`

const fireblockIDPub = `-----BEGIN PGP PUBLIC KEY BLOCK-----

mQENBFnXM/gBCACx0nazKmx52DIBJtEbNPGCYUvozxbNzz0u3O0TEqDuUwolzlBZ
ZwjmZ67T2gP16j87i7MNnVeJjQJyMnJ+O9U71dkNiSzAijDrDatK7QYcY/WUWp3C
cYOFCoFlsfCrgklf6qo2ZgGLputtpx+j59wK+S6ns2Wkt579q5feFJmbKxsspZ5x
E4E/JGFCEKyKzOU8NQo8E582iLtaBFyLU9ED5AO6iJ1ud9ehpxZv/9NMRW/rrDZ3
fzqFcgHXYURfZb303YOFJugazy5Fb8QelTZw00LzZX0OA/nvLPzy+fkKSmXlrW8l
pH8i/r/NI6jvXlnK6AhK9ZMCstYwpjMRfxQtABEBAAG0H0ZpcmVibG9jayBJZCA8
ZGV2QGZpcmVibG9jay5pbz6JATcEEwEIACEFAlnXM/gCGwMFCwkIBwIGFQgJCgsC
BBYCAwECHgECF4AACgkQyLyYjh8h63e/iAgAkecy+8MlbCwzbaJconaPL+hRcdpq
UY018ed9v2QADQ+xOZ4Mk7tZWlnGBmJ23uAwq+blMfN69IsmxH5WEqLgY+Rr/dTE
KkkT+VwhVEhjq+nlBN0quXURO/Wv0/SmJ6opml/4IXu9Ld5ksDpuH3dk/ZJve0Cw
a1MS1qWjM22x0zIwyPbGwCdErPmHDYvD5oOtIuwJzog60zrtGm1FUE+ZE0FltE6H
NsndsGQcvmTu+PVtXgOJgzy5vEHYD5MBuuyxtADi6CknRbtf/xd7nIGywjtQ8mJd
0M/7q407iEnE9+t9DOw4GLCPik6QBFsyspLFEuAgqEKst0idXBPTSt3smbkBDQRZ
1zP4AQgA8ZjCXX4xueH8AkQ8C0HnC9OuYCfIRlU/IvOqUlJ59ueqUROWUA4fAjdk
mvi2bdGykxEXMsOHST1c7J4mAnaBR68t61bc1GwzKeDTeKChq5eqiCV4+BJzbBA8
UAdtYkIypyG+XimG3nf4zwYiu5xsqXVuCeHsHf4jU8p1HL7kTl6Fs7QS9LJlUvt4
ZHpwIwyVW/0ObtXNqFjz+WgbjpbDgb469bYZiieIEoJofnjN3RUiNdKT8VWQQogK
dboMxe88oVcjV5MvjCYHKxpnC9c4uSsV3I9ZH7crg2zl+cLC7bvCaTVAogVEiorT
ln5khxQYpVgofMcGhiYCx15bVbmnawARAQABiQEfBBgBCAAJBQJZ1zP4AhsMAAoJ
EMi8mI4fIet3vgwH/1XEuZo2T3hPHQklahkgvbyVXNiCCX6Zx271iM456NIzagyP
kKqceBfDtByStODTXHhoJC3i+oP+s7r6R2XIddOLkN5d14D0egx422rmtHGyadfa
vbOhjTHQusFuRw+4foUCD6PFPTTk7VE7pVLdnWQhzhWMrctQMkV87BGumnS5xOuy
xmnjULYePDhsSf0017BXmxr0eIDk5BSon9EMK2kP5dElv/4yXPJaIAy0asDhrz8t
LSyKoYM8XjfB35uMeKdmgXFCSJo4rddJXiiejxoDpQvpAiVyVqzJEGWcnJId0inV
CGifs7Vr+vcnsX9qq49g8uJpQxCeD2sxSK2Hr8U=
=pNHs
-----END PGP PUBLIC KEY BLOCK-----`

const signatureClear2 = `-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA256

0xd2a84f4b8b650937ec8f73cd8be2c74add5a911ba64df27458ed8229da804a26
-----BEGIN PGP SIGNATURE-----

iQEbBAEBCAAGBQJZ1zjjAAoJEMi8mI4fIet3yQIH+JV+93bLVsB/TIKA6OENAQWq
OgWGof/SIqkqeB51uzUFPq4O9HU/DPOAXvVDzlFrL3Qe8qM3eOJSJNNZ74haMzoQ
Q+sjAne7LdFML97DOJklcdhQQoaIjCfo4DuwJiXM8yUEsqt1JHgur8v2V4Q9mPYf
kyWSkG0iv2xHYR92t7v3RvO0ogTjoM5pktcyFqHDTey9n032cbqOwzRoEN5c5Lx7
XFKbTdvqSAinR1gFoddh5HkfSmThIsM87CdFc9jQerY26+gT2s7hqhgj38gd4HGq
tSVc+uxzgAD07B9AxMMT/uN2RNwmCXwnKuOWKMOMwsU53r87QQ0V49Gjqnbd1A==
=YLYk
-----END PGP SIGNATURE-----`

func TestLoadKey(t *testing.T) {
	bFireblockPubKey := [][]byte{[]byte(fireblockPubKey)}
	_, error1 := readPGPKeys(bFireblockPubKey)
	assert.Equal(t, error1, nil, "no error when loading a PGP PubKey")
	bFireblockPrivKey := [][]byte{[]byte(fireblockPrivKey)}
	_, error2 := readPGPKeys(bFireblockPrivKey)
	assert.Equal(t, error2, nil, "no error when loading a PGP PrivKey")
	bSignatureClear := [][]byte{[]byte(signatureClear)}
	_, error3 := readPGPKeys(bSignatureClear)
	assert.NotEqual(t, error3, nil, "error when loading a signature")
}

func TestVerify(t *testing.T) {
	r, _ := PGPVerify([]byte(signatureClear2), [][]byte{[]byte(fireblockIDPub)})
	assert.Equal(t, r, true, "signature is valid")
}
