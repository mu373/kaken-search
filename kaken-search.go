package main

//
import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/jszwec/csvutil"
	"github.com/urfave/cli/v2"
)

type KakenResearcherFormatted struct {
	OriginalName           string
	OriginalRole           string
	Name                   string
	Keywords               string
	AffiliationInstitution string
	AffiliationDepartment  string
	AffiliationTitle       string
	KakenUrl               string
}

type KakenResearcher struct {
	TotalResults int `json:"totalResults"`
	StartIndex   int `json:"startIndex"`
	ItemsPerPage int `json:"itemsPerPage"`
	Researchers  []struct {
		AffiliationsCurrent []struct {
			Sequence int `json:"sequence"`
			Since    struct {
				CommonEraYear int `json:"commonEra:year"`
				Month         int `json:"month"`
				Day           int `json:"day"`
			} `json:"since"`
			Until struct {
				CommonEraYear int `json:"commonEra:year"`
				Month         int `json:"month"`
				Day           int `json:"day"`
			} `json:"until"`
			AffiliationInstitution struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
				IDInstitutionErad string `json:"id:institution:erad"`
			} `json:"affiliation:institution"`
			AffiliationDepartment struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
				IDDepartmentErad string `json:"id:department:erad"`
			} `json:"affiliation:department"`
			AffiliationJobTitle struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
				IDJobTitleErad string `json:"id:jobTitle:erad"`
			} `json:"affiliation:jobTitle"`
		} `json:"affiliations:current"`
		AffiliationsHistory []struct {
			Sequence int `json:"sequence"`
			Since    struct {
				CommonEraYear int `json:"commonEra:year"`
				Month         int `json:"month"`
				Day           int `json:"day"`
			} `json:"since"`
			Until struct {
				CommonEraYear int `json:"commonEra:year"`
				Month         int `json:"month"`
				Day           int `json:"day"`
			} `json:"until"`
			AffiliationInstitution struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
				IDInstitutionMext string `json:"id:institution:mext"`
			} `json:"affiliation:institution,omitempty"`
			AffiliationDepartment struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
				IDDepartmentMext string `json:"id:department:mext"`
				IDDepartmentErad string `json:"id:department:erad"`
			} `json:"affiliation:department,omitempty"`
			AffiliationJobTitle struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
				IDJobTitleMext string `json:"id:jobTitle:mext"`
				IDJobTitleErad string `json:"id:jobTitle:erad"`
			} `json:"affiliation:jobTitle,omitempty"`
			AffiliationDepartment0 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:department,omitempty"`
			AffiliationJobTitle0 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:jobTitle,omitempty"`
			AffiliationDepartment1 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
				IDDepartmentMext string `json:"id:department:mext"`
			} `json:"affiliation:department,omitempty"`
			AffiliationJobTitle1 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
				IDJobTitleMext string `json:"id:jobTitle:mext"`
			} `json:"affiliation:jobTitle,omitempty"`
			AffiliationInstitution0 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:institution,omitempty"`
			AffiliationDepartment2 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:department,omitempty"`
			AffiliationJobTitle2 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:jobTitle,omitempty"`
			AffiliationInstitution1 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:institution,omitempty"`
			AffiliationDepartment3 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:department,omitempty"`
			AffiliationJobTitle3 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:jobTitle,omitempty"`
			AffiliationInstitution2 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:institution,omitempty"`
			AffiliationDepartment4 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:department,omitempty"`
			AffiliationJobTitle4 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:jobTitle,omitempty"`
			AffiliationInstitution3 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:institution,omitempty"`
			AffiliationDepartment5 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:department,omitempty"`
			AffiliationJobTitle5 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:jobTitle,omitempty"`
			AffiliationInstitution4 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:institution,omitempty"`
			AffiliationDepartment6 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:department,omitempty"`
			AffiliationJobTitle6 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:jobTitle,omitempty"`
			AffiliationInstitution5 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:institution,omitempty"`
			AffiliationDepartment7 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:department,omitempty"`
			AffiliationJobTitle7 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:jobTitle,omitempty"`
			AffiliationInstitution6 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:institution,omitempty"`
			AffiliationDepartment8 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:department,omitempty"`
			AffiliationJobTitle8 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"affiliation:jobTitle,omitempty"`
		} `json:"affiliations:history"`
		RecordSource struct {
			IDPersonKakenhi []string `json:"id:person:kakenhi"`
		} `json:"recordSource"`
		EradPersonId            []string `json:"id:person:erad"`
		RelationRelatedResource []struct {
			IDProjectKakenhi string `json:"id:project:kakenhi"`
		} `json:"relation:relatedResource"`
		Name struct {
			Since struct {
				CommonEraYear int `json:"commonEra:year"`
				Month         int `json:"month"`
				Day           int `json:"day"`
			} `json:"since"`
			Until struct {
				CommonEraYear int `json:"commonEra:year"`
				Month         int `json:"month"`
				Day           int `json:"day"`
			} `json:"until"`
			HumanReadableValue []struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"humanReadableValue"`
			FamilyName []struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"name:familyName"`
			GivenName []struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"name:givenName"`
		} `json:"name"`
		Names []struct {
			Sequence           int `json:"sequence"`
			HumanReadableValue []struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"humanReadableValue"`
			FamilyName []struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"name:familyName"`
			GivenName []struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"name:givenName"`
			Since struct {
				CommonEraYear int `json:"commonEra:year"`
				Month         int `json:"month"`
				Day           int `json:"day"`
			} `json:"since"`
			Until struct {
				CommonEraYear int `json:"commonEra:year"`
				Month         int `json:"month"`
				Day           int `json:"day"`
			} `json:"until"`
		} `json:"names"`
		WorkProject []struct {
			RecordSource struct {
				IDProjectKakenhi []string `json:"id:project:kakenhi"`
			} `json:"recordSource"`
			Role []struct {
				RoleInProjectKakenhi string `json:"code:roleInProject:kakenhi"`
			} `json:"role"`
			ProjectStatus struct {
				StatusCode string `json:"statusCode"`
				FiscalYear struct {
					CommonEraYear  string `json:"commonEra:year"`
					FirstDateMonth int    `json:"firstDate:month"`
					FirstDateDay   int    `json:"firstDate:day"`
				} `json:"fiscal:year"`
			} `json:"projectStatus,omitempty"`
			Keyword []struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"keyword,omitempty"`
			Since struct {
				FiscalYear struct {
					CommonEraYear  string `json:"commonEra:year"`
					FirstDateMonth int    `json:"firstDate:month"`
					FirstDateDay   int    `json:"firstDate:day"`
				} `json:"fiscal:year"`
			} `json:"since"`
			Until struct {
				FiscalYear struct {
					CommonEraYear string `json:"commonEra:year"`
				} `json:"fiscal:year"`
			} `json:"until"`
			Title []struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"title"`
			Category []struct {
				HumanReadableValue []struct {
					Lang    string `json:"lang"`
					Path    string `json:"path"`
					NiiCode string `json:"niiCode"`
					Text    string `json:"text"`
				} `json:"humanReadableValue"`
			} `json:"category"`
			Field []struct {
				HumanReadableValue []struct {
					Lang       string `json:"lang"`
					Sequence   int    `json:"sequence"`
					Path       string `json:"path"`
					NiiCode    string `json:"niiCode"`
					Text       string `json:"text"`
					FieldTable string `json:"fieldTable"`
				} `json:"humanReadableValue"`
			} `json:"field,omitempty"`
			Institution []struct {
				HumanReadableValue []struct {
					Lang     string `json:"lang"`
					Sequence int    `json:"sequence"`
					NiiCode  string `json:"niiCode"`
					Text     string `json:"text"`
				} `json:"humanReadableValue"`
			} `json:"institution,omitempty"`
			Member []struct {
				Sequence int `json:"sequence"`
				Role     []struct {
					RoleInProjectKakenhi string `json:"code:roleInProject:kakenhi"`
				} `json:"role"`
				EradPersonId    string `json:"id:person:erad"`
				InstitutionName []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"institution:name"`
				DepartmentName []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"department:name"`
				JobTitle []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"jobTitle"`
				PersonName []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"person:name"`
			} `json:"member"`
			ReviewSection []struct {
				HumanReadableValue []struct {
					Lang      string `json:"lang"`
					Text      string `json:"text"`
					NiiCode   string `json:"niiCode"`
					TableType string `json:"tableType"`
					Sequence  int    `json:"sequence"`
				} `json:"humanReadableValue"`
			} `json:"review_section,omitempty"`
		} `json:"work:project"`
		WorkProduct []struct {
			Accn         []string `json:"accn"`
			RecordSource struct {
				IDProductKakenhi []string `json:"id:product:kakenhi"`
			} `json:"recordSource,omitempty"`
			RelationRelatedResource []struct {
				IDProjectKakenhi string `json:"id:project:kakenhi"`
			} `json:"relation:relatedResource"`
			ResourceType                  string `json:"resourceType"`
			DescriptionConferenceLocation struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"description:conferenceLocation,omitempty"`
			TitleMain struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"title:main"`
			DatePublicationDate struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
				Month         int `json:"month"`
				Day           int `json:"day"`
			} `json:"date:publicationDate,omitempty"`
			ContributorOrganizerUnparsed []struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"contributor:organizer:unparsed,omitempty"`
			CreatorUnparsed []struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"creator:unparsed,omitempty"`
			AttributeCreatorCandidate []struct {
				Estimated bool `json:"estimated"`
				List      []struct {
					Sequence   int `json:"sequence"`
					PersonName []struct {
						Text string `json:"text"`
						Lang string `json:"lang"`
					} `json:"person:name"`
					Accn []string `json:"accn"`
				} `json:"list"`
			} `json:"attribute:creator:candidate,omitempty"`
			AttributeInvited                    bool   `json:"attribute:invited,omitempty"`
			AttributeReviewed                   bool   `json:"attribute:reviewed,omitempty"`
			AttributeAcknowledgement            bool   `json:"attribute:acknowledgement,omitempty"`
			RightsAccessLevel                   string `json:"rights:accessLevel,omitempty"`
			AttributeJointInternationalResearch bool   `json:"attribute:jointInternationalResearch,omitempty"`
			SourceTitle                         struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"source:title,omitempty"`
			SourceFirstPage      string `json:"source:firstPage,omitempty"`
			SourceLastPage       string `json:"source:lastPage,omitempty"`
			SourceVolume         string `json:"source:volume,omitempty"`
			DatePublicationDate0 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate1 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate2 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate3 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate4 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate5 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate6 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate7 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate8 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			AttributeForeign     bool `json:"attribute:foreign,omitempty"`
			DatePublicationDate9 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DescriptionTotalPages int `json:"description:totalPages,omitempty"`
			DatePublicationDate10 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			ContributorPublisherUnparsed []struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"contributor:publisher:unparsed,omitempty"`
			DatePatentApplicationDate struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
				Month         int `json:"month"`
				Day           int `json:"day"`
			} `json:"date:patentApplicationDate,omitempty"`
			ContributorPatentInventorUnparsed []struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"contributor:patentInventor:unparsed,omitempty"`
			ContributorPatentAssigneeUnparsed []struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"contributor:patentAssignee:unparsed,omitempty"`
			DatePublicationDate11 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			IDNaid []struct {
				Reason string `json:"reason"`
				Value  string `json:"value"`
			} `json:"id:naid,omitempty"`
			DatePublicationDate12 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate13 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate14 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate15 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate16 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate17 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate18 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate19 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate20 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate21 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			IDDoi []struct {
				Value string `json:"value"`
			} `json:"id:doi,omitempty"`
			AttributeIDDoiAsWritten []struct {
				Text           string `json:"text"`
				ValueExtracted []struct {
					Value              string `json:"value"`
					ExistenceConfirmed bool   `json:"existenceConfirmed"`
				} `json:"valueExtracted"`
			} `json:"attribute:id:doi:asWritten,omitempty"`
			TitleAlternative []struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"title:alternative,omitempty"`
			DatePublicationDate22 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate23 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate24 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate25 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate26 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate27 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate28 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate29 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate30 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate31 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate32 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate33 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate34 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate35 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			IDPatentID struct {
				ApplicationType string `json:"applicationType"`
			} `json:"id:patentId,omitempty"`
			DatePublicationDate36 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate37 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate38 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate39 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate40 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate41 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate42 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate43 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate44 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate45 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate46 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate47 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate48 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate49 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate50 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate51 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate52 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate53 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate54 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate55 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate56 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate57 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate58 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate59 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate60 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate61 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate62 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate63 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			RecordSource0 struct {
				IDProductKakenhi []string `json:"id:product:kakenhi"`
				IDJalc           []string `json:"id:jalc"`
			} `json:"recordSource,omitempty"`
			Language               string `json:"language,omitempty"`
			SourceAlternativeTitle []struct {
				Text string `json:"text"`
				Lang string `json:"lang"`
			} `json:"source:alternativeTitle,omitempty"`
			SourceIssn []struct {
				Value string `json:"value"`
			} `json:"source:issn,omitempty"`
			SourceIssue           string `json:"source:issue,omitempty"`
			DatePublicationDate64 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate65 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate66 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate67 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate68 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate69 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate70 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate71 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate72 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate73 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate74 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate75 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate76 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate77 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate78 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate79 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate80 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate81 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate82 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate83 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate84 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate85 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate86 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate87 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate88 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate89 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate90 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate91 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate92 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate93 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate94 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate95 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate96 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate97 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate98 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate99 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate100 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate101 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate102 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate103 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate104 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate105 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate106 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate107 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate108 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate109 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate110 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate111 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate112 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate113 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate114 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate115 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate116 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate117 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate118 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate119 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate120 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate121 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate122 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate123 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate124 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate125 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate126 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate127 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate128 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate129 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate130 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate131 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate132 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate133 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate134 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate135 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate136 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate137 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate138 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate139 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate140 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate141 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			RecordSource1 struct {
				IDProductKakenhi []string `json:"id:product:kakenhi"`
				IDJalc           []string `json:"id:jalc"`
			} `json:"recordSource,omitempty"`
			DatePublicationDate142 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate143 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate144 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate145 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate146 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate147 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			IDPatentID0 struct {
				ApplicationNumber []string `json:"applicationNumber"`
			} `json:"id:patentId,omitempty"`
			DatePublicationDate148 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate149 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate150 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate151 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate152 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate153 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate154 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate155 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate156 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate157 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate158 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate159 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			RecordSource2 struct {
				IDProductKakenhi []string `json:"id:product:kakenhi"`
				IDJalc           []string `json:"id:jalc"`
			} `json:"recordSource,omitempty"`
			DatePublicationDate160 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate161 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate162 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate163 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate164 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate165 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate166 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate167 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate168 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate169 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate170 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate171 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate172 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate173 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate174 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate175 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate176 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate177 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate178 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate179 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			RecordSource3 struct {
				IDProductKakenhi []string `json:"id:product:kakenhi"`
				IDJalc           []string `json:"id:jalc"`
			} `json:"recordSource,omitempty"`
			DatePublicationDate180 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate181 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate182 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate183 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate184 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate185 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate186 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate187 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate188 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate189 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate190 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate191 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate192 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			IDPatentID1 struct {
				ApplicationNumber []string `json:"applicationNumber"`
			} `json:"id:patentId,omitempty"`
			DatePublicationDate193 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate194 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate195 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate196 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate197 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate198 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate199 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate200 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate201 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate202 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate203 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate204 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate205 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			RecordSource4 struct {
				IDProductKakenhi []string `json:"id:product:kakenhi"`
				IDJalc           []string `json:"id:jalc"`
			} `json:"recordSource,omitempty"`
			DatePublicationDate206 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate207 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate208 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate209 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate210 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate211 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate212 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate213 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate214 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate215 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate216 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate217 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate218 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate219 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate220 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate221 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate222 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate223 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate224 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate225 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate226 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate227 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate228 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate229 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate230 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate231 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate232 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate233 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate234 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate235 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate236 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate237 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate238 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate239 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate240 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate241 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate242 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate243 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate244 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate245 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate246 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate247 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate248 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate249 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate250 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate251 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate252 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate253 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate254 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate255 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate256 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			RecordSource5 struct {
				IDProductKakenhi []string `json:"id:product:kakenhi"`
				IDJalc           []string `json:"id:jalc"`
			} `json:"recordSource,omitempty"`
			DatePublicationDate257 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate258 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate259 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate260 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate261 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate262 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate263 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate264 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate265 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate266 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate267 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate268 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate269 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate270 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate271 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate272 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate273 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate274 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate275 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate276 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate277 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate278 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate279 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate280 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate281 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate282 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate283 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate284 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate285 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate286 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate287 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate288 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			RecordSource6 struct {
				IDProductKakenhi []string `json:"id:product:kakenhi"`
				IDJalc           []string `json:"id:jalc"`
			} `json:"recordSource,omitempty"`
			DatePublicationDate289 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate290 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate291 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate292 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate293 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate294 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate295 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate296 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate297 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate298 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate299 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate300 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate301 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate302 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate303 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate304 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
			DatePublicationDate305 struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
				} `json:"humanReadableValue"`
				CommonEraYear int `json:"commonEra:year"`
			} `json:"date:publicationDate,omitempty"`
		} `json:"work:product"`
		Ongoing struct {
			Title []struct {
				HumanReadableValue []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"humanReadableValue"`
			} `json:"title"`
			Role []struct {
				RoleInProjectKakenhi string `json:"code:roleInProject:kakenhi"`
			} `json:"role"`
		} `json:"ongoing"`
		Accn string `json:"accn"`
	} `json:"researchers"`
}

func main() {
	var kakenAppId string
	var researcherName string
	var researcherAffiliation string
	var outputFormat string

	var inputPath string
	var outputPath string
	var nameColumnIndex int
	var affiliationColumnIndex int

	app := &cli.App{
		Name:  "kaken-search",
		Usage: "a CLI tool to search for researchers on KAKEN database",
		Commands: []*cli.Command{
			{
				Name:    "single",
				Aliases: []string{"s"},
				Usage:   "Search a single researcher",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "id",
						Value:       "",
						Usage:       "CiNii App ID",
						Destination: &kakenAppId,
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "name",
						Aliases:     []string{"n"},
						Value:       "",
						Usage:       "Researcher's name",
						Destination: &researcherName,
					},
					&cli.StringFlag{
						Name:        "affl",
						Aliases:     []string{"a"},
						Value:       "",
						Usage:       "Researcher's affiliation (institution)",
						Destination: &researcherAffiliation,
					},
					&cli.StringFlag{
						Name:        "format",
						Aliases:     []string{"f"},
						Value:       "",
						Usage:       "Output format [json]",
						Destination: &outputFormat,
					},
				},
				Action: func(cCtx *cli.Context) error {
					SingleSearch(kakenAppId, researcherName, researcherAffiliation, outputFormat)
					return nil
				},
			},
			{
				Name:    "bulk",
				Aliases: []string{"b"},
				Usage:   "Search multiple researchers from CSV",
				Action: func(cCtx *cli.Context) error {
					BulkSearch(kakenAppId, inputPath, outputPath, nameColumnIndex, affiliationColumnIndex)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "id",
						Value:       "",
						Usage:       "CiNii App ID",
						Destination: &kakenAppId,
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "input",
						Aliases:     []string{"i"},
						Value:       "",
						Usage:       "Path for input CSV",
						Destination: &inputPath,
						Required:    true,
					},
					&cli.IntFlag{
						Name:        "name",
						Aliases:     []string{"n"},
						Value:       0,
						Usage:       "Column number containing researcher's name (start counting from 0)",
						Destination: &nameColumnIndex,
					},
					&cli.IntFlag{
						Name:        "affl",
						Aliases:     []string{"a"},
						Value:       1,
						Usage:       "Column number containing researcher's affiliation (start counting from 0)",
						Destination: &affiliationColumnIndex,
					},
					&cli.StringFlag{
						Name:        "output",
						Aliases:     []string{"o"},
						Value:       "output.tsv",
						Usage:       "Path for output TSV",
						Destination: &outputPath,
					},
				},
			},
		},
		Flags: []cli.Flag{},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func SearchAndFormat(kakenAppId string, researcherName string, researcherRole string) (KakenResearcherFormatted, error) {

	researcherData, err := GetResearcherDataFromAPI(researcherName, researcherRole, kakenAppId)
	if err != nil {
		fmt.Printf(err.Error())
	}

	formattedData := FormatResearcherData(researcherData)
	formattedData.OriginalName = researcherName
	formattedData.OriginalRole = researcherRole

	return formattedData, err
}

func SingleSearch(kakenAppId string, researcherName string, researcherRole string, outputFormat string) {
	formattedData, _ := SearchAndFormat(kakenAppId, researcherName, researcherRole)

	if outputFormat == "json" {
		jsonData, err := json.Marshal(formattedData)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(string(jsonData))
	} else {
		fmt.Println(researcherName + "：" + researcherRole)
		fmt.Println(formattedData.Keywords)
	}
}

func BulkSearch(kakenAppId string, inputPath string, outputPath string, nameColumnIndex int, affiliationColumnIndex int) {

	// 名前と所属を含んだCSVを読み込み
	filePath := inputPath
	csvRow, err := ReadCSVColumns(filePath, []int{nameColumnIndex, affiliationColumnIndex}, true)
	if err != nil {
		fmt.Printf("Failed to read reseacher names from CSV: %s\n", err)
		return
	}

	// 検索結果書き込み用TSVの準備
	f, err := os.Create(outputPath)
	if err != nil {
		fmt.Println(err)
	}
	w := csv.NewWriter(f)
	w.Comma = '\t'
	enc := csvutil.NewEncoder(w)

	var researcherName string
	var researcherRole string
	for _, row := range csvRow {

		if len(row) >= 2 {
			researcherName = row[0]
			researcherRole = row[1]
		} else {
			researcherName = row[0]
			researcherRole = ""
		}

		fmt.Println(researcherName + "：" + researcherRole)
		formattedData, _ := SearchAndFormat(kakenAppId, researcherName, researcherRole)
		fmt.Println(formattedData.Keywords)

		if err := enc.Encode(formattedData); err != nil {
			fmt.Println(err)
		}

		time.Sleep(500 * time.Millisecond)
	}

	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Println(err)
	}

}

func GetResearcherDataFromAPI(researcherName string, researcherInstitution string, kakenAppId string) (KakenResearcher, error) {
	// API Endpoint
	apiUrl := &url.URL{}
	apiUrl.Scheme = "https"
	apiUrl.Host = "nrid.nii.ac.jp"
	apiUrl.Path = "opensearch"

	apiQuery := apiUrl.Query()
	apiQuery.Set("appid", kakenAppId)
	apiQuery.Set("qg", researcherName)

	if researcherInstitution != "" {
		apiQuery.Set("kw", researcherInstitution)
	}
	apiQuery.Set("format", "json")
	apiUrl.RawQuery = apiQuery.Encode()

	// HTTPSでJSONデータを取得
	response, err := http.Get(fmt.Sprint(apiUrl))
	if err != nil {
		return KakenResearcher{}, fmt.Errorf("Failed to fetch data: %w", err)
		// return
	}
	defer response.Body.Close()

	// レスポンスボディを読み込む
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return KakenResearcher{}, fmt.Errorf("Failed to read response body: %s\n", err)
	}

	// JSONデコード
	var data KakenResearcher
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(body)
		return KakenResearcher{}, fmt.Errorf("Failed to decode JSON: %s\n", err)
	}

	return data, nil
}

func FormatResearcherData(data KakenResearcher) KakenResearcherFormatted {

	var newData KakenResearcherFormatted

	// 検索結果があった場合
	if data.TotalResults != 0 {
		researcher := data.Researchers[0]

		newData.Name = researcher.Name.HumanReadableValue[0].Text

		if len(researcher.AffiliationsCurrent) != 0 {
			newData.AffiliationInstitution = researcher.AffiliationsCurrent[0].AffiliationInstitution.HumanReadableValue[0].Text
			newData.AffiliationDepartment = researcher.AffiliationsCurrent[0].AffiliationDepartment.HumanReadableValue[0].Text
			newData.AffiliationTitle = researcher.AffiliationsCurrent[0].AffiliationJobTitle.HumanReadableValue[0].Text
		}

		eradId := researcher.EradPersonId[0]
		newData.KakenUrl = fmt.Sprintf("https://nrid.nii.ac.jp/ja/nrid/10000%s/", eradId)

		// 研究代表者である研究課題からキーワードを抽出
		var keywordTexts []string
		projects := researcher.WorkProject
		for _, project := range projects {
			if project.Role[0].RoleInProjectKakenhi == "principal_investigator" {
				keywords := project.Keyword
				for _, keyword := range keywords {
					keywordText := keyword.HumanReadableValue[0].Text
					keywordTexts = append(keywordTexts, keywordText)
				}
			}
		}
		newData.Keywords = strings.Join(CreateUniqueSlice(keywordTexts), ",")
	}

	return newData
}

func CreateUniqueSlice(slice []string) []string {
	type keyValue struct {
		Key   string
		Value int
	}

	countMap := make(map[string]int)
	for _, value := range slice {
		countMap[value]++
	}

	var keyValueList []keyValue
	for key, value := range countMap {
		keyValueList = append(keyValueList, keyValue{Key: key, Value: value})
	}

	sort.Slice(keyValueList, func(i, j int) bool {
		return keyValueList[i].Value > keyValueList[j].Value
	})

	uniqueSlice := make([]string, 0)
	uniqueMap := make(map[string]bool)
	for _, kv := range keyValueList {
		if !uniqueMap[kv.Key] {
			uniqueSlice = append(uniqueSlice, kv.Key)
			uniqueMap[kv.Key] = true
		}
	}

	return uniqueSlice
}

func ReadCSVColumns(filePath string, columnIndices []int, hasHeader bool) ([][]string, error) {
	// CSVファイルを開く
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// CSVリーダーを作成
	reader := csv.NewReader(file)

	// ヘッダー行をスキップ
	if hasHeader {
		_, err = reader.Read()
		if err != nil {
			return nil, err
		}
	}

	// 列を読み込み
	var rows [][]string
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		var row []string
		for _, columnIndex := range columnIndices {
			if columnIndex < len(record) {
				row = append(row, record[columnIndex])
			}
		}
		rows = append(rows, row)
	}

	return rows, nil
}
