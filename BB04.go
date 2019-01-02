


package main

import (
"crypto/sha256"
"fmt"
"github.com/Nik-U/pbc"
)

type  Reg  struct{
	pairing  *pbc.Pairing
	g1,g2  *pbc.Element
	pub *Reg_pubkey
	pri *Reg_prikey
}
type  Reg_sig  struct{
	delta,r  *pbc.Element
}
type  Reg_pubkey  struct{
	u,v  *pbc.Element
}
type  Reg_prikey struct {
	x, y *pbc.Element
}
type  Sys  struct{
	g1,g2 *pbc.Element
	pairing *pbc.Pairing
}

func  Sys_Start()(*Sys){
	params:=pbc.GenerateA(160,512)
	pairing:=params.NewPairing()
	g1:=pairing.NewG1().Rand()
	g2:=pairing.NewG2().Rand()
	system:=new (Sys)
	system.g2=g2
	system.g1=g1
	system.pairing=pairing
	return  system


}

func  (sys  *Sys)BB04_KenGen()(*pbc.Element,*pbc.Element,*pbc.Element,*pbc.Element,*Reg){
	x:=sys.pairing.NewZr().Rand()
	y:=sys.pairing.NewZr().Rand()
	u:=sys.pairing.NewG2().PowZn(sys.g2,x)
	v:=sys.pairing.NewG2().PowZn(sys.g2,y)
	Reg:=new(Reg)
	Reg.pub=new(Reg_pubkey)
	Reg.pri=new(Reg_prikey)
	Reg.pairing=sys.pairing
	Reg.g1=sys.g1
	Reg.g2=sys.g2
	Reg.pri.x=x
	Reg.pri.y=y
	Reg.pub.u=u
	Reg.pub.v=v
	return x,y,u,v,Reg
}


func  (Reg  *Reg)BB04_Sign(msg string)(*pbc.Element,*pbc.Element){
	MsgHash:=Reg.pairing.NewZr().SetFromStringHash(msg,sha256.New())
	r:=Reg.pairing.NewZr().Rand()
	xxx:=Reg.pairing.NewZr().Add(Reg.pri.x,MsgHash)
	yyy:=Reg.pairing.NewZr().Mul(Reg.pri.y,r)
	zzz:=Reg.pairing.NewZr().Add(xxx,yyy)
	uuu:=Reg.pairing.NewZr().Invert(zzz)
	delta:=Reg.pairing.NewG1().PowZn(Reg.g1,uuu)
	return   delta,r
}
func  (Reg  *Reg)BB04_Ver(msg  string,r,delta *pbc.Element)bool{
	MsgHash:=Reg.pairing.NewZr().SetFromStringHash(msg,sha256.New())
	yyy:=Reg.pairing.NewG2().PowZn(Reg.g2,MsgHash)
	xxx:=Reg.pairing.NewG2().PowZn(Reg.pub.v,r)
	zzz:=Reg.pairing.NewG2().Mul(yyy,xxx)
	uuu:=Reg.pairing.NewG2().Mul(zzz,Reg.pub.u)
	e1:=Reg.pairing.NewGT().Pair(delta,uuu)
	e2:=Reg.pairing.NewGT().Pair(Reg.g1,Reg.g2)
	if !e1.Equals(e2) {
		fmt.Println("*BUG* Org_Signature check failed *BUG*")
		return false
	} else {
		fmt.Println("Org_Signature verified correctly")
		return  true
	}
}
