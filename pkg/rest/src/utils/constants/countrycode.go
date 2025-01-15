package constants

import (
	"encoding/json"

	"github.com/petmeds24/backend/pkg/rest/src/models"
)

func CountryData() ([]models.Countries, error) {
	var countries []models.Countries
	err := json.Unmarshal([]byte(data), &countries)
	if err != nil {
		return nil, err
	}
	return countries, nil
}

const data = `[
    {
        "country": "Afghanistan",
        "country_code": "AF",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/AF.svg",
        "country_phone_code": "+93",
        "unicode": "U+1F1E6 U+1F1EB"
    },
    {
        "country": "Albania",
        "country_code": "AL",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/AL.svg",
        "country_phone_code": "+355",
        "unicode": "U+1F1E6 U+1F1F1"
    },
    {
        "country": "Andorra",
        "country_code": "AD",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/AD.svg",
        "country_phone_code": "+376",
        "unicode": "U+1F1E6 U+1F1E9"
    },
    {
        "country": "Angola",
        "country_code": "AO",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/AO.svg",
        "country_phone_code": "+244",
        "unicode": "U+1F1E6 U+1F1F4"
    },
    {
        "country": "Anguilla",
        "country_code": "AI",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/AI.svg",
        "country_phone_code": "+1264",
        "unicode": "U+1F1E6 U+1F1EE"
    },
    {
        "country": "Antarctica",
        "country_code": "AQ",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/AQ.svg",
        "country_phone_code": "+672",
        "unicode": "U+1F1E6 U+1F1F6"
    },
    {
        "country": "Antigua & Barbuda",
        "country_code": "AG",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/AG.svg",
        "country_phone_code": "+1268",
        "unicode": "U+1F1E6 U+1F1EC"
    },
    {
        "country": "Argentina",
        "country_code": "AR",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/AR.svg",
        "country_phone_code": "+54",
        "unicode": "U+1F1E6 U+1F1F7"
    },
    {
        "country": "Armenia",
        "country_code": "AM",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/AM.svg",
        "country_phone_code": "+374",
        "unicode": "U+1F1E6 U+1F1F2"
    },
    {
        "country": "Aruba",
        "country_code": "AW",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/AW.svg",
        "country_phone_code": "+297",
        "unicode": "U+1F1E6 U+1F1FC"
    },
    {
        "country": "Australia",
        "country_code": "AU",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/AU.svg",
        "country_phone_code": "+61",
        "unicode": "U+1F1E6 U+1F1FA"
    },
    {
        "country": "Austria",
        "country_code": "AT",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/AT.svg",
        "country_phone_code": "+43",
        "unicode": "U+1F1E6 U+1F1F9"
    },
    {
        "country": "Azerbaijan",
        "country_code": "AZ",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/AZ.svg",
        "country_phone_code": "+994",
        "unicode": "U+1F1E6 U+1F1FF"
    },
    {
        "country": "Bahamas",
        "country_code": "BS",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BS.svg",
        "country_phone_code": "+1242",
        "unicode": "U+1F1E7 U+1F1F8"
    },
    {
        "country": "Bahrain",
        "country_code": "BH",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BH.svg",
        "country_phone_code": "+973",
        "unicode": "U+1F1E7 U+1F1ED"
    },
    {
        "country": "Bangladesh",
        "country_code": "BD",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BD.svg",
        "country_phone_code": "+880",
        "unicode": "U+1F1E7 U+1F1E9"
    },
    {
        "country": "Barbados",
        "country_code": "BB",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BB.svg",
        "country_phone_code": "+1246",
        "unicode": "U+1F1E7 U+1F1E7"
    },
    {
        "country": "Belarus",
        "country_code": "BY",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BY.svg",
        "country_phone_code": "+375",
        "unicode": "U+1F1E7 U+1F1FE"
    },
    {
        "country": "Belgium",
        "country_code": "BE",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BE.svg",
        "country_phone_code": "+32",
        "unicode": "U+1F1E7 U+1F1EA"
    },
    {
        "country": "Belize",
        "country_code": "BZ",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BZ.svg",
        "country_phone_code": "+501",
        "unicode": "U+1F1E7 U+1F1FF"
    },
    {
        "country": "Benin",
        "country_code": "BJ",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BJ.svg",
        "country_phone_code": "+229",
        "unicode": "U+1F1E7 U+1F1EF"
    },
    {
        "country": "Bermuda",
        "country_code": "BM",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BM.svg",
        "country_phone_code": "+1441",
        "unicode": "U+1F1E7 U+1F1F2"
    },
    {
        "country": "Bhutan",
        "country_code": "BT",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BT.svg",
        "country_phone_code": "+975",
        "unicode": "U+1F1E7 U+1F1F9"
    },
    {
        "country": "Bolivia",
        "country_code": "BO",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BO.svg",
        "country_phone_code": "+591",
        "unicode": "U+1F1E7 U+1F1F4"
    },
    {
        "country": "Bosnia & Herzegovina",
        "country_code": "BA",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BA.svg",
        "country_phone_code": "+387",
        "unicode": "U+1F1E7 U+1F1E6"
    },
    {
        "country": "Botswana",
        "country_code": "BW",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BW.svg",
        "country_phone_code": "+267",
        "unicode": "U+1F1E7 U+1F1FC"
    },
    {
        "country": "Brazil",
        "country_code": "BR",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BR.svg",
        "country_phone_code": "+55",
        "unicode": "U+1F1E7 U+1F1F7"
    },
    {
        "country": "Brunei",
        "country_code": "BN",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BN.svg",
        "country_phone_code": "+673",
        "unicode": "U+1F1E7 U+1F1F3"
    },
    {
        "country": "Bulgaria",
        "country_code": "BG",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BG.svg",
        "country_phone_code": "+359",
        "unicode": "U+1F1E7 U+1F1EC"
    },
    {
        "country": "Burkina Faso",
        "country_code": "BF",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BF.svg",
        "country_phone_code": "+226",
        "unicode": "U+1F1E7 U+1F1EB"
    },
    {
        "country": "Burundi",
        "country_code": "BI",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/BI.svg",
        "country_phone_code": "+257",
        "unicode": "U+1F1E7 U+1F1EE"
    },
    {
        "country": "Cabo Verde",
        "country_code": "CV",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/CV.svg",
        "country_phone_code": "+238",
        "unicode": "U+1F1E8 U+1F1FB"
    },
    {
        "country": "Cambodia",
        "country_code": "KH",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/KH.svg",
        "country_phone_code": "+855",
        "unicode": "U+1F1F0 U+1F1ED"
    },
    {
        "country": "Cameroon",
        "country_code": "CM",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/CM.svg",
        "country_phone_code": "+237",
        "unicode": "U+1F1E8 U+1F1F2"
    },
    {
        "country": "Canada",
        "country_code": "CA",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/CA.svg",
        "country_phone_code": "+1",
        "unicode": "U+1F1E8 U+1F1E6"
    },
    {
        "country": "Cayman Islands",
        "country_code": "KY",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/KY.svg",
        "country_phone_code": "+1345",
        "unicode": "U+1F1F0 U+1F1FE"
    },
    {
        "country": "Central African Republic",
        "country_code": "CF",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/CF.svg",
        "country_phone_code": "+236",
        "unicode": "U+1F1E8 U+1F1EB"
    },
    {
        "country": "Chad",
        "country_code": "TD",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/TD.svg",
        "country_phone_code": "+235",
        "unicode": "U+1F1F9 U+1F1E9"
    },
    {
        "country": "Chile",
        "country_code": "CL",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/CL.svg",
        "country_phone_code": "+56",
        "unicode": "U+1F1E8 U+1F1F1"
    },
    {
        "country": "China",
        "country_code": "CN",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/CN.svg",
        "country_phone_code": "+86",
        "unicode": "U+1F1E8 U+1F1F3"
    },
    {
        "country": "Colombia",
        "country_code": "CO",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/CO.svg",
        "country_phone_code": "+57",
        "unicode": "U+1F1E8 U+1F1F4"
    },
    {
        "country": "Comoros",
        "country_code": "KM",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/KM.svg",
        "country_phone_code": "+269",
        "unicode": "U+1F1F0 U+1F1F2"
    },
    {
        "country": "Congo - Brazzaville",
        "country_code": "CG",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/CG.svg",
        "country_phone_code": "+242",
        "unicode": "U+1F1E8 U+1F1EC"
    },
    {
        "country": "Congo - Kinshasa",
        "country_code": "CD",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/CD.svg",
        "country_phone_code": "+243",
        "unicode": "U+1F1E8 U+1F1E9"
    },
    {
        "country": "Cook Islands",
        "country_code": "CK",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/CK.svg",
        "country_phone_code": "+682",
        "unicode": "U+1F1E8 U+1F1F0"
    },
    {
        "country": "Costa Rica",
        "country_code": "CR",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/CR.svg",
        "country_phone_code": "+506",
        "unicode": "U+1F1E8 U+1F1F7"
    },
    {
        "country": "Croatia",
        "country_code": "HR",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/HR.svg",
        "country_phone_code": "+385",
        "unicode": "U+1F1ED U+1F1F7"
    },
    {
        "country": "Cuba",
        "country_code": "CU",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/CU.svg",
        "country_phone_code": "+53",
        "unicode": "U+1F1E8 U+1F1FA"
    },
    {
        "country": "Cura\u00e7ao",
        "country_code": "CW",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/CW.svg",
        "country_phone_code": "+599",
        "unicode": "U+1F1E8 U+1F1FC"
    },
    {
        "country": "Cyprus",
        "country_code": "CY",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/CY.svg",
        "country_phone_code": "+357",
        "unicode": "U+1F1E8 U+1F1FE"
    },
    {
        "country": "Czechia",
        "country_code": "CZ",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/CZ.svg",
        "country_phone_code": "+420",
        "unicode": "U+1F1E8 U+1F1FF"
    },
    {
        "country": "C\u00f4te d\u2019Ivoire",
        "country_code": "CI",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/CI.svg",
        "country_phone_code": "+225",
        "unicode": "U+1F1E8 U+1F1EE"
    },
    {
        "country": "Denmark",
        "country_code": "DK",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/DK.svg",
        "country_phone_code": "+45",
        "unicode": "U+1F1E9 U+1F1F0"
    },
    {
        "country": "Djibouti",
        "country_code": "DJ",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/DJ.svg",
        "country_phone_code": "+253",
        "unicode": "U+1F1E9 U+1F1EF"
    },
    {
        "country": "Dominica",
        "country_code": "DM",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/DM.svg",
        "country_phone_code": "+1767",
        "unicode": "U+1F1E9 U+1F1F2"
    },
    {
        "country": "Dominican Republic",
        "country_code": "DO",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/DO.svg",
        "country_phone_code": "+1 809, +1 829, +1 849",
        "unicode": "U+1F1E9 U+1F1F4"
    },
    {
        "country": "Ecuador",
        "country_code": "EC",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/EC.svg",
        "country_phone_code": "+593",
        "unicode": "U+1F1EA U+1F1E8"
    },
    {
        "country": "Egypt",
        "country_code": "EG",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/EG.svg",
        "country_phone_code": "+20",
        "unicode": "U+1F1EA U+1F1EC"
    },
    {
        "country": "El Salvador",
        "country_code": "SV",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SV.svg",
        "country_phone_code": "+503",
        "unicode": "U+1F1F8 U+1F1FB"
    },
    {
        "country": "Equatorial Guinea",
        "country_code": "GQ",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GQ.svg",
        "country_phone_code": "+240",
        "unicode": "U+1F1EC U+1F1F6"
    },
    {
        "country": "Eritrea",
        "country_code": "ER",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/ER.svg",
        "country_phone_code": "+291",
        "unicode": "U+1F1EA U+1F1F7"
    },
    {
        "country": "Estonia",
        "country_code": "EE",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/EE.svg",
        "country_phone_code": "+372",
        "unicode": "U+1F1EA U+1F1EA"
    },
    {
        "country": "Eswatini",
        "country_code": "SZ",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SZ.svg",
        "country_phone_code": "+268",
        "unicode": "U+1F1F8 U+1F1FF"
    },
    {
        "country": "Ethiopia",
        "country_code": "ET",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/ET.svg",
        "country_phone_code": "+251",
        "unicode": "U+1F1EA U+1F1F9"
    },
    {
        "country": "Falkland Islands",
        "country_code": "FK",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/FK.svg",
        "country_phone_code": "+500",
        "unicode": "U+1F1EB U+1F1F0"
    },
    {
        "country": "Faroe Islands",
        "country_code": "FO",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/FO.svg",
        "country_phone_code": "+298",
        "unicode": "U+1F1EB U+1F1F4"
    },
    {
        "country": "Fiji",
        "country_code": "FJ",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/FJ.svg",
        "country_phone_code": "+679",
        "unicode": "U+1F1EB U+1F1EF"
    },
    {
        "country": "Finland",
        "country_code": "FI",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/FI.svg",
        "country_phone_code": "+358",
        "unicode": "U+1F1EB U+1F1EE"
    },
    {
        "country": "France",
        "country_code": "FR",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/FR.svg",
        "country_phone_code": "+33",
        "unicode": "U+1F1EB U+1F1F7"
    },
    {
        "country": "French Guiana",
        "country_code": "GF",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GF.svg",
        "country_phone_code": "+594",
        "unicode": "U+1F1EC U+1F1EB"
    },
    {
        "country": "French Polynesia",
        "country_code": "PF",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/PF.svg",
        "country_phone_code": "+689",
        "unicode": "U+1F1F5 U+1F1EB"
    },
    {
        "country": "Gabon",
        "country_code": "GA",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GA.svg",
        "country_phone_code": "+241",
        "unicode": "U+1F1EC U+1F1E6"
    },
    {
        "country": "Gambia",
        "country_code": "GM",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GM.svg",
        "country_phone_code": "+220",
        "unicode": "U+1F1EC U+1F1F2"
    },
    {
        "country": "Georgia",
        "country_code": "GE",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GE.svg",
        "country_phone_code": "+995",
        "unicode": "U+1F1EC U+1F1EA"
    },
    {
        "country": "Germany",
        "country_code": "DE",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/DE.svg",
        "country_phone_code": "+49",
        "unicode": "U+1F1E9 U+1F1EA"
    },
    {
        "country": "Ghana",
        "country_code": "GH",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GH.svg",
        "country_phone_code": "+233",
        "unicode": "U+1F1EC U+1F1ED"
    },
    {
        "country": "Gibraltar",
        "country_code": "GI",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GI.svg",
        "country_phone_code": "+350",
        "unicode": "U+1F1EC U+1F1EE"
    },
    {
        "country": "Greece",
        "country_code": "GR",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GR.svg",
        "country_phone_code": "+30",
        "unicode": "U+1F1EC U+1F1F7"
    },
    {
        "country": "Greenland",
        "country_code": "GL",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GL.svg",
        "country_phone_code": "+299",
        "unicode": "U+1F1EC U+1F1F1"
    },
    {
        "country": "Grenada",
        "country_code": "GD",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GD.svg",
        "country_phone_code": "+1473",
        "unicode": "U+1F1EC U+1F1E9"
    },
    {
        "country": "Guadeloupe",
        "country_code": "GP",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GP.svg",
        "country_phone_code": "+590",
        "unicode": "U+1F1EC U+1F1F5"
    },
    {
        "country": "Guam",
        "country_code": "GU",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GU.svg",
        "country_phone_code": "+1671",
        "unicode": "U+1F1EC U+1F1FA"
    },
    {
        "country": "Guatemala",
        "country_code": "GT",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GT.svg",
        "country_phone_code": "+502",
        "unicode": "U+1F1EC U+1F1F9"
    },
    {
        "country": "Guernsey",
        "country_code": "GG",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GG.svg",
        "country_phone_code": "+44",
        "unicode": "U+1F1EC U+1F1EC"
    },
    {
        "country": "Guinea",
        "country_code": "GN",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GN.svg",
        "country_phone_code": "+224",
        "unicode": "U+1F1EC U+1F1F3"
    },
    {
        "country": "Guinea-Bissau",
        "country_code": "GW",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GW.svg",
        "country_phone_code": "+245",
        "unicode": "U+1F1EC U+1F1FC"
    },
    {
        "country": "Guyana",
        "country_code": "GY",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GY.svg",
        "country_phone_code": "+592",
        "unicode": "U+1F1EC U+1F1FE"
    },
    {
        "country": "Haiti",
        "country_code": "HT",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/HT.svg",
        "country_phone_code": "+509",
        "unicode": "U+1F1ED U+1F1F9"
    },
    {
        "country": "Honduras",
        "country_code": "HN",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/HN.svg",
        "country_phone_code": "+504",
        "unicode": "U+1F1ED U+1F1F3"
    },
    {
        "country": "Hong Kong SAR China",
        "country_code": "HK",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/HK.svg",
        "country_phone_code": "+852",
        "unicode": "U+1F1ED U+1F1F0"
    },
    {
        "country": "Hungary",
        "country_code": "HU",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/HU.svg",
        "country_phone_code": "+36",
        "unicode": "U+1F1ED U+1F1FA"
    },
    {
        "country": "Iceland",
        "country_code": "IS",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/IS.svg",
        "country_phone_code": "+354",
        "unicode": "U+1F1EE U+1F1F8"
    },
    {
        "country": "India",
        "country_code": "IN",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/IN.svg",
        "country_phone_code": "+91",
        "unicode": "U+1F1EE U+1F1F3"
    },
    {
        "country": "Indonesia",
        "country_code": "ID",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/ID.svg",
        "country_phone_code": "+62",
        "unicode": "U+1F1EE U+1F1E9"
    },
    {
        "country": "Iran",
        "country_code": "IR",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/IR.svg",
        "country_phone_code": "+98",
        "unicode": "U+1F1EE U+1F1F7"
    },
    {
        "country": "Iraq",
        "country_code": "IQ",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/IQ.svg",
        "country_phone_code": "+964",
        "unicode": "U+1F1EE U+1F1F6"
    },
    {
        "country": "Ireland",
        "country_code": "IE",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/IE.svg",
        "country_phone_code": "+353",
        "unicode": "U+1F1EE U+1F1EA"
    },
    {
        "country": "Isle of Man",
        "country_code": "IM",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/IM.svg",
        "country_phone_code": "+44",
        "unicode": "U+1F1EE U+1F1F2"
    },
    {
        "country": "Israel",
        "country_code": "IL",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/IL.svg",
        "country_phone_code": "+972",
        "unicode": "U+1F1EE U+1F1F1"
    },
    {
        "country": "Italy",
        "country_code": "IT",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/IT.svg",
        "country_phone_code": "+39",
        "unicode": "U+1F1EE U+1F1F9"
    },
    {
        "country": "Jamaica",
        "country_code": "JM",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/JM.svg",
        "country_phone_code": "+1 876",
        "unicode": "U+1F1EF U+1F1F2"
    },
    {
        "country": "Japan",
        "country_code": "JP",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/JP.svg",
        "country_phone_code": "+81",
        "unicode": "U+1F1EF U+1F1F5"
    },
    {
        "country": "Jersey",
        "country_code": "JE",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/JE.svg",
        "country_phone_code": "+44",
        "unicode": "U+1F1EF U+1F1EA"
    },
    {
        "country": "Jordan",
        "country_code": "JO",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/JO.svg",
        "country_phone_code": "+962",
        "unicode": "U+1F1EF U+1F1F4"
    },
    {
        "country": "Kazakhstan",
        "country_code": "KZ",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/KZ.svg",
        "country_phone_code": "+7",
        "unicode": "U+1F1F0 U+1F1FF"
    },
    {
        "country": "Kenya",
        "country_code": "KE",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/KE.svg",
        "country_phone_code": "+254",
        "unicode": "U+1F1F0 U+1F1EA"
    },
    {
        "country": "Kiribati",
        "country_code": "KI",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/KI.svg",
        "country_phone_code": "+686",
        "unicode": "U+1F1F0 U+1F1EE"
    },
    {
        "country": "Kosovo",
        "country_code": "XK",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/XK.svg",
        "country_phone_code": "+383",
        "unicode": "U+1F1FD U+1F1F0"
    },
    {
        "country": "Kuwait",
        "country_code": "KW",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/KW.svg",
        "country_phone_code": "+965",
        "unicode": "U+1F1F0 U+1F1FC"
    },
    {
        "country": "Kyrgyzstan",
        "country_code": "KG",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/KG.svg",
        "country_phone_code": "+996",
        "unicode": "U+1F1F0 U+1F1EC"
    },
    {
        "country": "Laos",
        "country_code": "LA",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/LA.svg",
        "country_phone_code": "+856",
        "unicode": "U+1F1F1 U+1F1E6"
    },
    {
        "country": "Latvia",
        "country_code": "LV",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/LV.svg",
        "country_phone_code": "+371",
        "unicode": "U+1F1F1 U+1F1FB"
    },
    {
        "country": "Lebanon",
        "country_code": "LB",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/LB.svg",
        "country_phone_code": "+961",
        "unicode": "U+1F1F1 U+1F1E7"
    },
    {
        "country": "Lesotho",
        "country_code": "LS",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/LS.svg",
        "country_phone_code": "+266",
        "unicode": "U+1F1F1 U+1F1F8"
    },
    {
        "country": "Liberia",
        "country_code": "LR",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/LR.svg",
        "country_phone_code": "+231",
        "unicode": "U+1F1F1 U+1F1F7"
    },
    {
        "country": "Libya",
        "country_code": "LY",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/LY.svg",
        "country_phone_code": "+218",
        "unicode": "U+1F1F1 U+1F1FE"
    },
    {
        "country": "Liechtenstein",
        "country_code": "LI",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/LI.svg",
        "country_phone_code": "+423",
        "unicode": "U+1F1F1 U+1F1EE"
    },
    {
        "country": "Lithuania",
        "country_code": "LT",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/LT.svg",
        "country_phone_code": "+370",
        "unicode": "U+1F1F1 U+1F1F9"
    },
    {
        "country": "Luxembourg",
        "country_code": "LU",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/LU.svg",
        "country_phone_code": "+352",
        "unicode": "U+1F1F1 U+1F1FA"
    },
    {
        "country": "Macao SAR China",
        "country_code": "MO",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MO.svg",
        "country_phone_code": "+853",
        "unicode": "U+1F1F2 U+1F1F4"
    },
    {
        "country": "Madagascar",
        "country_code": "MG",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MG.svg",
        "country_phone_code": "+261",
        "unicode": "U+1F1F2 U+1F1EC"
    },
    {
        "country": "Malawi",
        "country_code": "MW",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MW.svg",
        "country_phone_code": "+265",
        "unicode": "U+1F1F2 U+1F1FC"
    },
    {
        "country": "Malaysia",
        "country_code": "MY",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MY.svg",
        "country_phone_code": "+60",
        "unicode": "U+1F1F2 U+1F1FE"
    },
    {
        "country": "Maldives",
        "country_code": "MV",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MV.svg",
        "country_phone_code": "+960",
        "unicode": "U+1F1F2 U+1F1FB"
    },
    {
        "country": "Mali",
        "country_code": "ML",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/ML.svg",
        "country_phone_code": "+223",
        "unicode": "U+1F1F2 U+1F1F1"
    },
    {
        "country": "Malta",
        "country_code": "MT",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MT.svg",
        "country_phone_code": "+356",
        "unicode": "U+1F1F2 U+1F1F9"
    },
    {
        "country": "Marshall Islands",
        "country_code": "MH",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MH.svg",
        "country_phone_code": "+692",
        "unicode": "U+1F1F2 U+1F1ED"
    },
    {
        "country": "Martinique",
        "country_code": "MQ",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MQ.svg",
        "country_phone_code": "+596",
        "unicode": "U+1F1F2 U+1F1F6"
    },
    {
        "country": "Mauritania",
        "country_code": "MR",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MR.svg",
        "country_phone_code": "+222",
        "unicode": "U+1F1F2 U+1F1F7"
    },
    {
        "country": "Mauritius",
        "country_code": "MU",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MU.svg",
        "country_phone_code": "+230",
        "unicode": "U+1F1F2 U+1F1FA"
    },
    {
        "country": "Mayotte",
        "country_code": "YT",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/YT.svg",
        "country_phone_code": "+262",
        "unicode": "U+1F1FE U+1F1F9"
    },
    {
        "country": "Mexico",
        "country_code": "MX",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MX.svg",
        "country_phone_code": "+52",
        "unicode": "U+1F1F2 U+1F1FD"
    },
    {
        "country": "Micronesia",
        "country_code": "FM",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/FM.svg",
        "country_phone_code": "+691",
        "unicode": "U+1F1EB U+1F1F2"
    },
    {
        "country": "Moldova",
        "country_code": "MD",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MD.svg",
        "country_phone_code": "+373",
        "unicode": "U+1F1F2 U+1F1E9"
    },
    {
        "country": "Monaco",
        "country_code": "MC",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MC.svg",
        "country_phone_code": "+377",
        "unicode": "U+1F1F2 U+1F1E8"
    },
    {
        "country": "Mongolia",
        "country_code": "MN",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MN.svg",
        "country_phone_code": "+976",
        "unicode": "U+1F1F2 U+1F1F3"
    },
    {
        "country": "Montenegro",
        "country_code": "ME",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/ME.svg",
        "country_phone_code": "+382",
        "unicode": "U+1F1F2 U+1F1EA"
    },
    {
        "country": "Montserrat",
        "country_code": "MS",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MS.svg",
        "country_phone_code": "+1 664",
        "unicode": "U+1F1F2 U+1F1F8"
    },
    {
        "country": "Morocco",
        "country_code": "MA",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MA.svg",
        "country_phone_code": "+212",
        "unicode": "U+1F1F2 U+1F1E6"
    },
    {
        "country": "Mozambique",
        "country_code": "MZ",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MZ.svg",
        "country_phone_code": "+258",
        "unicode": "U+1F1F2 U+1F1FF"
    },
    {
        "country": "Myanmar (Burma)",
        "country_code": "MM",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MM.svg",
        "country_phone_code": "+95",
        "unicode": "U+1F1F2 U+1F1F2"
    },
    {
        "country": "Namibia",
        "country_code": "NA",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/NA.svg",
        "country_phone_code": "+264",
        "unicode": "U+1F1F3 U+1F1E6"
    },
    {
        "country": "Nauru",
        "country_code": "NR",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/NR.svg",
        "country_phone_code": "+674",
        "unicode": "U+1F1F3 U+1F1F7"
    },
    {
        "country": "Nepal",
        "country_code": "NP",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/NP.svg",
        "country_phone_code": "+977",
        "unicode": "U+1F1F3 U+1F1F5"
    },
    {
        "country": "Netherlands",
        "country_code": "NL",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/NL.svg",
        "country_phone_code": "+31",
        "unicode": "U+1F1F3 U+1F1F1"
    },
    {
        "country": "New Caledonia",
        "country_code": "NC",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/NC.svg",
        "country_phone_code": "+687",
        "unicode": "U+1F1F3 U+1F1E8"
    },
    {
        "country": "New Zealand",
        "country_code": "NZ",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/NZ.svg",
        "country_phone_code": "+64",
        "unicode": "U+1F1F3 U+1F1FF"
    },
    {
        "country": "Nicaragua",
        "country_code": "NI",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/NI.svg",
        "country_phone_code": "+505",
        "unicode": "U+1F1F3 U+1F1EE"
    },
    {
        "country": "Niger",
        "country_code": "NE",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/NE.svg",
        "country_phone_code": "+227",
        "unicode": "U+1F1F3 U+1F1EA"
    },
    {
        "country": "Nigeria",
        "country_code": "NG",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/NG.svg",
        "country_phone_code": "+234",
        "unicode": "U+1F1F3 U+1F1EC"
    },
    {
        "country": "Niue",
        "country_code": "NU",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/NU.svg",
        "country_phone_code": "+683",
        "unicode": "U+1F1F3 U+1F1FA"
    },
    {
        "country": "Norfolk Island",
        "country_code": "NF",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/NF.svg",
        "country_phone_code": "+672",
        "unicode": "U+1F1F3 U+1F1EB"
    },
    {
        "country": "North Korea",
        "country_code": "KP",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/KP.svg",
        "country_phone_code": "+850",
        "unicode": "U+1F1F0 U+1F1F5"
    },
    {
        "country": "Northern Mariana Islands",
        "country_code": "MP",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/MP.svg",
        "country_phone_code": "+1 670",
        "unicode": "U+1F1F2 U+1F1F5"
    },
    {
        "country": "Norway",
        "country_code": "NO",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/NO.svg",
        "country_phone_code": "+47",
        "unicode": "U+1F1F3 U+1F1F4"
    },
    {
        "country": "Oman",
        "country_code": "OM",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/OM.svg",
        "country_phone_code": "+968",
        "unicode": "U+1F1F4 U+1F1F2"
    },
    {
        "country": "Pakistan",
        "country_code": "PK",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/PK.svg",
        "country_phone_code": "+92",
        "unicode": "U+1F1F5 U+1F1F0"
    },
    {
        "country": "Palau",
        "country_code": "PW",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/PW.svg",
        "country_phone_code": "+680",
        "unicode": "U+1F1F5 U+1F1FC"
    },
    {
        "country": "Palestinian Territories",
        "country_code": "PS",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/PS.svg",
        "country_phone_code": "+970",
        "unicode": "U+1F1F5 U+1F1F8"
    },
    {
        "country": "Panama",
        "country_code": "PA",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/PA.svg",
        "country_phone_code": "+507",
        "unicode": "U+1F1F5 U+1F1E6"
    },
    {
        "country": "Papua New Guinea",
        "country_code": "PG",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/PG.svg",
        "country_phone_code": "+675",
        "unicode": "U+1F1F5 U+1F1EC"
    },
    {
        "country": "Paraguay",
        "country_code": "PY",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/PY.svg",
        "country_phone_code": "+595",
        "unicode": "U+1F1F5 U+1F1FE"
    },
    {
        "country": "Peru",
        "country_code": "PE",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/PE.svg",
        "country_phone_code": "+51",
        "unicode": "U+1F1F5 U+1F1EA"
    },
    {
        "country": "Philippines",
        "country_code": "PH",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/PH.svg",
        "country_phone_code": "+63",
        "unicode": "U+1F1F5 U+1F1ED"
    },
    {
        "country": "Pitcairn Islands",
        "country_code": "PN",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/PN.svg",
        "country_phone_code": "+64",
        "unicode": "U+1F1F5 U+1F1F3"
    },
    {
        "country": "Poland",
        "country_code": "PL",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/PL.svg",
        "country_phone_code": "+48",
        "unicode": "U+1F1F5 U+1F1F1"
    },
    {
        "country": "Portugal",
        "country_code": "PT",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/PT.svg",
        "country_phone_code": "+351",
        "unicode": "U+1F1F5 U+1F1F9"
    },
    {
        "country": "Puerto Rico",
        "country_code": "PR",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/PR.svg",
        "country_phone_code": "+1 787",
        "unicode": "U+1F1F5 U+1F1F7"
    },
    {
        "country": "Qatar",
        "country_code": "QA",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/QA.svg",
        "country_phone_code": "+974",
        "unicode": "U+1F1F6 U+1F1E6"
    },
    {
        "country": "Romania",
        "country_code": "RO",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/RO.svg",
        "country_phone_code": "+40",
        "unicode": "U+1F1F7 U+1F1F4"
    },
    {
        "country": "Russia",
        "country_code": "RU",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/RU.svg",
        "country_phone_code": "+7",
        "unicode": "U+1F1F7 U+1F1FA"
    },
    {
        "country": "Rwanda",
        "country_code": "RW",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/RW.svg",
        "country_phone_code": "+250",
        "unicode": "U+1F1F7 U+1F1FC"
    },
    {
        "country": "R\u00e9union",
        "country_code": "RE",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/RE.svg",
        "country_phone_code": "+262",
        "unicode": "U+1F1F7 U+1F1EA"
    },
    {
        "country": "Saint Helena",
        "country_code": "SH",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SH.svg",
        "country_phone_code": "+290",
        "unicode": "U+1F1F8 U+1F1ED"
    },
    {
        "country": "Saint Kitts & Nevis",
        "country_code": "KN",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/KN.svg",
        "country_phone_code": "+1 869",
        "unicode": "U+1F1F0 U+1F1F3"
    },
    {
        "country": "Saint Lucia",
        "country_code": "LC",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/LC.svg",
        "country_phone_code": "+1 758",
        "unicode": "U+1F1F1 U+1F1E8"
    },
    {
        "country": "Saint Pierre & Miquelon",
        "country_code": "PM",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/PM.svg",
        "country_phone_code": "+508",
        "unicode": "U+1F1F5 U+1F1F2"
    },
    {
        "country": "Samoa",
        "country_code": "WS",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/WS.svg",
        "country_phone_code": "+685",
        "unicode": "U+1F1FC U+1F1F8"
    },
    {
        "country": "San Marino",
        "country_code": "SM",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SM.svg",
        "country_phone_code": "+378",
        "unicode": "U+1F1F8 U+1F1F2"
    },
    {
        "country": "Sao Tome & Principe",
        "country_code": "ST",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/ST.svg",
        "country_phone_code": "+239",
        "unicode": "U+1F1F8 U+1F1F9"
    },
    {
        "country": "Saudi Arabia",
        "country_code": "SA",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SA.svg",
        "country_phone_code": "+966",
        "unicode": "U+1F1F8 U+1F1E6"
    },
    {
        "country": "Senegal",
        "country_code": "SN",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SN.svg",
        "country_phone_code": "+221",
        "unicode": "U+1F1F8 U+1F1F3"
    },
    {
        "country": "Serbia",
        "country_code": "RS",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/RS.svg",
        "country_phone_code": "+381",
        "unicode": "U+1F1F7 U+1F1F8"
    },
    {
        "country": "Seychelles",
        "country_code": "SC",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SC.svg",
        "country_phone_code": "+248",
        "unicode": "U+1F1F8 U+1F1E8"
    },
    {
        "country": "Sierra Leone",
        "country_code": "SL",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SL.svg",
        "country_phone_code": "+232",
        "unicode": "U+1F1F8 U+1F1F1"
    },
    {
        "country": "Singapore",
        "country_code": "SG",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SG.svg",
        "country_phone_code": "+65",
        "unicode": "U+1F1F8 U+1F1EC"
    },
    {
        "country": "Sint Maarten",
        "country_code": "SX",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SX.svg",
        "country_phone_code": "+1 721",
        "unicode": "U+1F1F8 U+1F1FD"
    },
    {
        "country": "Slovakia",
        "country_code": "SK",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SK.svg",
        "country_phone_code": "+421",
        "unicode": "U+1F1F8 U+1F1F0"
    },
    {
        "country": "Slovenia",
        "country_code": "SI",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SI.svg",
        "country_phone_code": "+386",
        "unicode": "U+1F1F8 U+1F1EE"
    },
    {
        "country": "Solomon Islands",
        "country_code": "SB",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SB.svg",
        "country_phone_code": "+677",
        "unicode": "U+1F1F8 U+1F1E7"
    },
    {
        "country": "Somalia",
        "country_code": "SO",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SO.svg",
        "country_phone_code": "+252",
        "unicode": "U+1F1F8 U+1F1F4"
    },
    {
        "country": "South Africa",
        "country_code": "ZA",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/ZA.svg",
        "country_phone_code": "+27",
        "unicode": "U+1F1FF U+1F1E6"
    },
    {
        "country": "South Korea",
        "country_code": "KR",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/KR.svg",
        "country_phone_code": "+82",
        "unicode": "U+1F1F0 U+1F1F7"
    },
    {
        "country": "South Sudan",
        "country_code": "SS",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SS.svg",
        "country_phone_code": "+211",
        "unicode": "U+1F1F8 U+1F1F8"
    },
    {
        "country": "Spain",
        "country_code": "ES",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/ES.svg",
        "country_phone_code": "+34",
        "unicode": "U+1F1EA U+1F1F8"
    },
    {
        "country": "Sri Lanka",
        "country_code": "LK",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/LK.svg",
        "country_phone_code": "+94",
        "unicode": "U+1F1F1 U+1F1F0"
    },
    {
        "country": "Sudan",
        "country_code": "SD",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SD.svg",
        "country_phone_code": "+249",
        "unicode": "U+1F1F8 U+1F1E9"
    },
    {
        "country": "Suriname",
        "country_code": "SR",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SR.svg",
        "country_phone_code": "+597",
        "unicode": "U+1F1F8 U+1F1F7"
    },
    {
        "country": "Sweden",
        "country_code": "SE",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SE.svg",
        "country_phone_code": "+46",
        "unicode": "U+1F1F8 U+1F1EA"
    },
    {
        "country": "Switzerland",
        "country_code": "CH",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/CH.svg",
        "country_phone_code": "+41",
        "unicode": "U+1F1E8 U+1F1ED"
    },
    {
        "country": "Syria",
        "country_code": "SY",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/SY.svg",
        "country_phone_code": "+963",
        "unicode": "U+1F1F8 U+1F1FE"
    },
    {
        "country": "Taiwan",
        "country_code": "TW",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/TW.svg",
        "country_phone_code": "+886",
        "unicode": "U+1F1F9 U+1F1FC"
    },
    {
        "country": "Tajikistan",
        "country_code": "TJ",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/TJ.svg",
        "country_phone_code": "+992",
        "unicode": "U+1F1F9 U+1F1EF"
    },
    {
        "country": "Tanzania",
        "country_code": "TZ",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/TZ.svg",
        "country_phone_code": "+255",
        "unicode": "U+1F1F9 U+1F1FF"
    },
    {
        "country": "Thailand",
        "country_code": "TH",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/TH.svg",
        "country_phone_code": "+66",
        "unicode": "U+1F1F9 U+1F1ED"
    },
    {
        "country": "Timor-Leste",
        "country_code": "TL",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/TL.svg",
        "country_phone_code": "+670",
        "unicode": "U+1F1F9 U+1F1F1"
    },
    {
        "country": "Togo",
        "country_code": "TG",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/TG.svg",
        "country_phone_code": "+228",
        "unicode": "U+1F1F9 U+1F1EC"
    },
    {
        "country": "Tokelau",
        "country_code": "TK",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/TK.svg",
        "country_phone_code": "+690",
        "unicode": "U+1F1F9 U+1F1F0"
    },
    {
        "country": "Tonga",
        "country_code": "TO",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/TO.svg",
        "country_phone_code": "+676",
        "unicode": "U+1F1F9 U+1F1F4"
    },
    {
        "country": "Trinidad & Tobago",
        "country_code": "TT",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/TT.svg",
        "country_phone_code": "+1 868",
        "unicode": "U+1F1F9 U+1F1F9"
    },
    {
        "country": "Tunisia",
        "country_code": "TN",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/TN.svg",
        "country_phone_code": "+216",
        "unicode": "U+1F1F9 U+1F1F3"
    },
    {
        "country": "Turkey",
        "country_code": "TR",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/TR.svg",
        "country_phone_code": "+90",
        "unicode": "U+1F1F9 U+1F1F7"
    },
    {
        "country": "Turkmenistan",
        "country_code": "TM",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/TM.svg",
        "country_phone_code": "+993",
        "unicode": "U+1F1F9 U+1F1F2"
    },
    {
        "country": "Turks & Caicos Islands",
        "country_code": "TC",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/TC.svg",
        "country_phone_code": "+1 649",
        "unicode": "U+1F1F9 U+1F1E8"
    },
    {
        "country": "Tuvalu",
        "country_code": "TV",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/TV.svg",
        "country_phone_code": "+688",
        "unicode": "U+1F1F9 U+1F1FB"
    },
    {
        "country": "Uganda",
        "country_code": "UG",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/UG.svg",
        "country_phone_code": "+256",
        "unicode": "U+1F1FA U+1F1EC"
    },
    {
        "country": "Ukraine",
        "country_code": "UA",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/UA.svg",
        "country_phone_code": "+380",
        "unicode": "U+1F1FA U+1F1E6"
    },
    {
        "country": "United Arab Emirates",
        "country_code": "AE",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/AE.svg",
        "country_phone_code": "+971",
        "unicode": "U+1F1E6 U+1F1EA"
    },
    {
        "country": "United Kingdom",
        "country_code": "GB",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/GB.svg",
        "country_phone_code": "+44",
        "unicode": "U+1F1EC U+1F1E7"
    },
    {
        "country": "United States",
        "country_code": "US",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/US.svg",
        "country_phone_code": "+1",
        "unicode": "U+1F1FA U+1F1F8"
    },
    {
        "country": "Uruguay",
        "country_code": "UY",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/UY.svg",
        "country_phone_code": "+598",
        "unicode": "U+1F1FA U+1F1FE"
    },
    {
        "country": "Uzbekistan",
        "country_code": "UZ",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/UZ.svg",
        "country_phone_code": "+998",
        "unicode": "U+1F1FA U+1F1FF"
    },
    {
        "country": "Vanuatu",
        "country_code": "VU",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/VU.svg",
        "country_phone_code": "+678",
        "unicode": "U+1F1FB U+1F1FA"
    },
    {
        "country": "Vatican City",
        "country_code": "VA",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/VA.svg",
        "country_phone_code": "+379",
        "unicode": "U+1F1FB U+1F1E6"
    },
    {
        "country": "Venezuela",
        "country_code": "VE",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/VE.svg",
        "country_phone_code": "+58",
        "unicode": "U+1F1FB U+1F1EA"
    },
    {
        "country": "Vietnam",
        "country_code": "VN",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/VN.svg",
        "country_phone_code": "+84",
        "unicode": "U+1F1FB U+1F1F3"
    },
    {
        "country": "Wallis & Futuna",
        "country_code": "WF",
        "country_image_url": "https://cdn.jsdelivr.net/npm/country-flag-emoji-json@2.0.0/dist/images/WF.svg",
        "country_phone_code": "+681",
        "unicode": "U+1F1FC U+1F1EB"
    }
]`
