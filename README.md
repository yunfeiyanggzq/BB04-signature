# Discription
it  is  a  BB04 signature  lib in golang 
# How  to  install  the  lib
## First :install  the  GMP
This package must be compiled using cgo. It also requires the installation of GMP and PBC. During the build process, this package will attempt to include <gmp.h> and <pbc/pbc.h>, and then dynamically link to GMP and PBC.

Most systems include a package for GMP. To install GMP in Debian / Ubuntu:

`sudo apt-get install libgmp-dev`

For an RPM installation with YUM:

`sudo yum install gmp-devel`

For installation with Fink (http://www.finkproject.org/) on Mac OS X:

`sudo fink install gmp gmp-shlibs`

For more information or to compile from source, visit https://gmplib.org/

## Second:install the PBC 
To install the PBC library, download the appropriate files for your system from https://crypto.stanford.edu/pbc/download.html. PBC has three dependencies: the gcc compiler, flex (http://flex.sourceforge.net/), and bison (https://www.gnu.org/software/bison/). See the respective sites for installation instructions. Most distributions include packages for these libraries. For example, in Debian / Ubuntu:

`sudo apt-get install build-essential flex bison`

The PBC source can be compiled and installed using the usual GNU Build System:
```
./configure
make
sudo make install
```

After installing, you may need to rebuild the search path for libraries:

`sudo ldconfig`

It is possible to install the package on Windows through the use of MinGW and MSYS. MSYS is required for installing PBC, while GMP can be installed through a package. Based on your MinGW installation, you may need to add` "-I/usr/local/include"` to CPPFLAGS and `"-L/usr/local/lib" `to LDFLAGS when building PBC. Likewise, you may need to add these options to` CGO_CPPFLAGS `and` CGO_LDFLAGS` when installing this package. 

and  then  install the  golang  pbc  lib 

` go  get  github.com/Nik-U/pbc`

## Third: install  the BB04 signature  golang  lib
download the bb04  signature  lib  

` go  get  go  get  github.com/yunfeiyangbuaa/BB04-signature`

and  imprt in your code 

`import "github.com/yunfeiyangbuaa/BB04-signature"`
#  How to use
##   Exmple
 ```
func  main(){
	msg:="hello  world"
	sys:=Sys_Start()
	_,_,_,_,prikey:=sys.BB04_KenGen()
	delta,r:=prikey.BB04_Sign(msg)
	prikey.BB04_Ver(msg,r,delta)
}
 ```
## Function 
```
func (Reg *Reg)BB04_Ver(msg string,r,delta *pbc.Element)bool
func (Reg *Reg)BB04_Sign(msg string)(*pbc.Element,*pbc.Element)
func (sys *Sys)BB04_KenGen()(*pbc.Element,*pbc.Element,*pbc.Element,*pbc.Element,*Reg)
func Sys_Start()(*Sys)
 ```
