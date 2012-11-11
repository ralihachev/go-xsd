//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		schemas.opengis.net/kml/2.2.0/atom-author-link.xsd
package gopkg_SchemasOpengisNetKml220AtomAuthorLinkXsd

//	There is no official atom XSD. This XSD is created based on:
//	http://atompub.org/2005/08/17/atom.rnc. A subset of Atom as used in the
//	ogckml22.xsd is defined here.
import (
	xsdt "github.com/metaleap/go-xsd/types"
)

type XsdGoPkgHasCdata struct { CombinedCharDatas string `xml:",chardata"` }

type TatomMediaType xsdt.String

//	Since TatomMediaType is just a simple String type, this merely sets the current value from the specified string.
func (me *TatomMediaType) SetFromString (s string) { (*xsdt.String)(me).SetFromString(s) }

//	Since TatomMediaType is just a simple String type, this merely returns the current string value.
func (me TatomMediaType) String () string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TatomMediaType's alias type xsdt.String
func (me TatomMediaType) ToXsdtString () xsdt.String { return xsdt.String(me) }

type TatomLanguageTag xsdt.String

//	Since TatomLanguageTag is just a simple String type, this merely sets the current value from the specified string.
func (me *TatomLanguageTag) SetFromString (s string) { (*xsdt.String)(me).SetFromString(s) }

//	Since TatomLanguageTag is just a simple String type, this merely returns the current string value.
func (me TatomLanguageTag) String () string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TatomLanguageTag's alias type xsdt.String
func (me TatomLanguageTag) ToXsdtString () xsdt.String { return xsdt.String(me) }

type TatomEmailAddress xsdt.String

//	Since TatomEmailAddress is just a simple String type, this merely sets the current value from the specified string.
func (me *TatomEmailAddress) SetFromString (s string) { (*xsdt.String)(me).SetFromString(s) }

//	Since TatomEmailAddress is just a simple String type, this merely returns the current string value.
func (me TatomEmailAddress) String () string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TatomEmailAddress's alias type xsdt.String
func (me TatomEmailAddress) ToXsdtString () xsdt.String { return xsdt.String(me) }

type TatomPersonConstruct struct {
	XsdGoPkgHasElems_Name

	XsdGoPkgHasElems_Uri

	XsdGoPkgHasElems_Email

}

type XsdGoPkgHasElems_Name struct {
	Names []xsdt.String `xml:"http://www.w3.org/2005/Atom name"`
}

type XsdGoPkgHasElem_Name struct {
	Name xsdt.String `xml:"http://www.w3.org/2005/Atom name"`
}

type XsdGoPkgHasElem_Uri struct {
	Uri xsdt.String `xml:"http://www.w3.org/2005/Atom uri"`
}

type XsdGoPkgHasElems_Uri struct {
	Uris []xsdt.String `xml:"http://www.w3.org/2005/Atom uri"`
}

type XsdGoPkgHasElem_Email struct {
	Email TatomEmailAddress `xml:"http://www.w3.org/2005/Atom email"`
}

type XsdGoPkgHasElems_Email struct {
	Emails []TatomEmailAddress `xml:"http://www.w3.org/2005/Atom email"`
}

type XsdGoPkgHasElems_Author struct {
	Authors []*TatomPersonConstruct `xml:"http://www.w3.org/2005/Atom author"`
}

type XsdGoPkgHasElem_Author struct {
	Author *TatomPersonConstruct `xml:"http://www.w3.org/2005/Atom author"`
}

type XsdGoPkgHasAttr_Href_XsdtString_ struct {
	Href xsdt.String `xml:"http://www.w3.org/2005/Atom href,attr"`
}

type XsdGoPkgHasAttr_Rel_XsdtString_ struct {
	Rel xsdt.String `xml:"http://www.w3.org/2005/Atom rel,attr"`
}

type XsdGoPkgHasAttr_Type_TatomMediaType_ struct {
	Type TatomMediaType `xml:"http://www.w3.org/2005/Atom type,attr"`
}

type XsdGoPkgHasAttr_Hreflang_TatomLanguageTag_ struct {
	Hreflang TatomLanguageTag `xml:"http://www.w3.org/2005/Atom hreflang,attr"`
}

type XsdGoPkgHasAttr_Title_XsdtString_ struct {
	Title xsdt.String `xml:"http://www.w3.org/2005/Atom title,attr"`
}

type XsdGoPkgHasAttr_Length_XsdtString_ struct {
	Length xsdt.String `xml:"http://www.w3.org/2005/Atom length,attr"`
}

type TxsdLink struct {
	XsdGoPkgHasAttr_Type_TatomMediaType_

	XsdGoPkgHasAttr_Title_XsdtString_

	XsdGoPkgHasAttr_Length_XsdtString_

	XsdGoPkgHasAttr_Hreflang_TatomLanguageTag_

	XsdGoPkgHasAttr_Href_XsdtString_

	XsdGoPkgHasAttr_Rel_XsdtString_

}

type XsdGoPkgHasElem_Link struct {
	Link *TxsdLink `xml:"http://www.w3.org/2005/Atom link"`
}

type XsdGoPkgHasElems_Link struct {
	Links []*TxsdLink `xml:"http://www.w3.org/2005/Atom link"`
}