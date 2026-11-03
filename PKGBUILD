pkgname=curd
pkgver=1.1.4
pkgrel=3
pkgdesc="CLI anime player with AniList tracking and Anikoto provider"
arch=('x86_64')
url="https://github.com/itsmenewbie03/curd"
license=('GPL')
depends=('mpv' 'rofi' 'ueberzugpp')
makedepends=('go' 'git')
source=("curd::git+https://github.com/itsmenewbie03/curd.git#branch=feat/anikoto-provider")
sha256sums=('SKIP')

prepare() {
    cd "$srcdir/curd"
}

build() {
    cd "$srcdir/curd"
    go build -mod=vendor -trimpath -ldflags="-s -w" -o curd ./cmd/curd
}

package() {
    cd "$srcdir/curd"
    install -Dm755 curd "$pkgdir/usr/bin/curd"
}
