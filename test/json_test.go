package test

import (
	"encoding/json"
	"fmt"
	"gojson"
	"testing"
)


func Test_Validator(t *testing.T) {
	data := `{"id":5240423333333333333333,"name":"酷旅-mob-otv-2","male":true,"other":null}`
	err := gojson.CheckValid([]byte(data))
	if err != nil {
		fmt.Println(err)
	}
}

func Test_json(t *testing.T) {
	data := `{"id":524042,"name":"酷旅-mob-otv-2","male":true,"other":null}`
	object, err := gojson.FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
	} else {
		jsonBytes, _ := json.Marshal(object)
		fmt.Println(string(jsonBytes))
	}
}

func Test_json22(t *testing.T) {
	data := `{
	  "id": 524042,
	  "name": "酷旅-mob-otv-2",
	  "male": true,
	  "other": null
	}`
	object, err := gojson.FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
	} else {
		jsonBytes, _ := json.Marshal(object)
		fmt.Println(string(jsonBytes))
	}
}

func Test_json_object(t *testing.T) {
	data := `{
	  "id": -524042.5,
	  "name": "酷旅-mob-otv-2",
	  "male": true,
	  "children": {
		"id": -524042.5,
		"name": "酷旅-mob-otv-2",
		"male": true,
		"other": null
	  },
	  "other": null
	}`
	object, err := gojson.FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
	} else {
		jsonBytes, _ := json.Marshal(object)
		fmt.Println(string(jsonBytes))
	}
}

func Test_json_array(t *testing.T) {
	data := `{
	  "id": [
		-524042.5,
		231231.2
	  ],
	  "name": "酷旅-mob-otv-2",
	  "male": true,
	  "other": null
	}`
	object, err := gojson.FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
	} else {
		jsonBytes, _ := json.Marshal(object)
		fmt.Println(string(jsonBytes))
	}
}

func Test_json_array111(t *testing.T) {
	data := `{
	  "id": [
		-524042.5,
		231231444445555555
	  ],
	  "name": "酷旅-mob-otv-2",
	  "male": true,
	  "other": null
	}`
	object, err := gojson.FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(object.GetCoding("travel"))
	}
}



func Test_json_array_parse(t *testing.T) {
	data := `{
			  "id": [
				-524042.5,
				2312314444
			  ],
			  "name": "酷旅-mob-otv-2",
			  "male": true,
			  "other": null
			}`
	object, err := gojson.FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	idArray:=object.GetJsonObject("id").GetJsonArray()
	for i,v:=range idArray{
		fmt.Println(i,v.GetFloat64())
		fmt.Println(i,v.GetInt64())
	}
	name:=object.GetJsonObject("name").GetString()
	fmt.Println(name)
	male:=object.GetJsonObject("male").GetBool()
	fmt.Println(male)
	other:=object.GetJsonObject("other").GetInterface()
	fmt.Println(other)
}

func Test_json_array_parse111(t *testing.T) {
	data := `{
	"id": 1296269,
	"owner": {
		"login": "octocat",
		"id": 1,
		"avatar_url": "https://github.com/images/error/octocat_happy.gif",
		"gravatar_id": "somehexcode",
		"url": "https://api.github.com/users/octocat",
		"html_url": "https://github.com/octocat",
		"followers_url": "https://api.github.com/users/octocat/followers",
		"following_url": "https://api.github.com/users/octocat/following{/other_user}",
		"gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
		"organizations_url": "https://api.github.com/users/octocat/orgs",
		"repos_url": "https://api.github.com/users/octocat/repos",
		"events_url": "https://api.github.com/users/octocat/events{/privacy}",
		"received_events_url": "https://api.github.com/users/octocat/received_events",
		"type": "User",
		"site_admin": false
	},
	"name": "Hello-World",
	"full_name": "octocat/Hello-World",
	"description": "This your first repo!",
	"private": false,
	"fork": false,
	"url": "https://api.github.com/repos/octocat/Hello-World",
	"html_url": "https://github.com/octocat/Hello-World",
	"clone_url": "https://github.com/octocat/Hello-World.git",
	"git_url": "git://github.com/octocat/Hello-World.git",
	"ssh_url": "git@github.com:octocat/Hello-World.git",
	"svn_url": "https://svn.github.com/octocat/Hello-World",
	"mirror_url": "git://git.example.com/octocat/Hello-World",
	"homepage": "https://github.com",
	"language": null,
	"forks_count": 9,
	"stargazers_count": 80,
	"watchers_count": 80,
	"size": 108,
	"default_branch": "master",
	"open_issues_count": 0,
	"has_issues": true,
	"has_wiki": true,
	"has_downloads": true,
	"pushed_at": "2011-01-26T19:06:43Z",
	"created_at": "2011-01-26T19:01:12Z",
	"updated_at": "2011-01-26T19:14:43Z",
	"permissions": {
		"admin": false,
		"push": false,
		"pull": true
	},
	"subscribers_count": 42,
	"organization": {
		"login": "octocat",
		"id": 1,
		"avatar_url": "https://github.com/images/error/octocat_happy.gif",
		"gravatar_id": "somehexcode",
		"url": "https://api.github.com/users/octocat",
		"html_url": "https://github.com/octocat",
		"followers_url": "https://api.github.com/users/octocat/followers",
		"following_url": "https://api.github.com/users/octocat/following{/other_user}",
		"gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
		"organizations_url": "https://api.github.com/users/octocat/orgs",
		"repos_url": "https://api.github.com/users/octocat/repos",
		"events_url": "https://api.github.com/users/octocat/events{/privacy}",
		"received_events_url": "https://api.github.com/users/octocat/received_events",
		"type": "Organization",
		"site_admin": false
	},
	"parent": {
		"id": 1296269,
		"owner": {
			"login": "octocat",
			"id": 1,
			"avatar_url": "https://github.com/images/error/octocat_happy.gif",
			"gravatar_id": "somehexcode",
			"url": "https://api.github.com/users/octocat",
			"html_url": "https://github.com/octocat",
			"followers_url": "https://api.github.com/users/octocat/followers",
			"following_url": "https://api.github.com/users/octocat/following{/other_user}",
			"gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
			"organizations_url": "https://api.github.com/users/octocat/orgs",
			"repos_url": "https://api.github.com/users/octocat/repos",
			"events_url": "https://api.github.com/users/octocat/events{/privacy}",
			"received_events_url": "https://api.github.com/users/octocat/received_events",
			"type": "User",
			"site_admin": false
		},
		"name": "Hello-World",
		"full_name": "octocat/Hello-World",
		"description": "This your first repo!",
		"private": false,
		"fork": true,
		"url": "https://api.github.com/repos/octocat/Hello-World",
		"html_url": "https://github.com/octocat/Hello-World",
		"clone_url": "https://github.com/octocat/Hello-World.git",
		"git_url": "git://github.com/octocat/Hello-World.git",
		"ssh_url": "git@github.com:octocat/Hello-World.git",
		"svn_url": "https://svn.github.com/octocat/Hello-World",
		"mirror_url": "git://git.example.com/octocat/Hello-World",
		"homepage": "https://github.com",
		"language": null,
		"forks_count": 9,
		"stargazers_count": 80,
		"watchers_count": 80,
		"size": 108,
		"default_branch": "master",
		"open_issues_count": 0,
		"has_issues": true,
		"has_wiki": true,
		"has_downloads": true,
		"pushed_at": "2011-01-26T19:06:43Z",
		"created_at": "2011-01-26T19:01:12Z",
		"updated_at": "2011-01-26T19:14:43Z",
		"permissions": {
			"admin": false,
			"push": false,
			"pull": true
		}
	},
	"source": {
		"id": 1296269,
		"owner": {
			"login": "octocat",
			"id": 1,
			"avatar_url": "https://github.com/images/error/octocat_happy.gif",
			"gravatar_id": "somehexcode",
			"url": "https://api.github.com/users/octocat",
			"html_url": "https://github.com/octocat",
			"followers_url": "https://api.github.com/users/octocat/followers",
			"following_url": "https://api.github.com/users/octocat/following{/other_user}",
			"gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
			"organizations_url": "https://api.github.com/users/octocat/orgs",
			"repos_url": "https://api.github.com/users/octocat/repos",
			"events_url": "https://api.github.com/users/octocat/events{/privacy}",
			"received_events_url": "https://api.github.com/users/octocat/received_events",
			"type": "User",
			"site_admin": false
		},
		"name": "Hello-World",
		"full_name": "octocat/Hello-World",
		"description": "This your first repo!",
		"private": false,
		"fork": true,
		"url": "https://api.github.com/repos/octocat/Hello-World",
		"html_url": "https://github.com/octocat/Hello-World",
		"clone_url": "https://github.com/octocat/Hello-World.git",
		"git_url": "git://github.com/octocat/Hello-World.git",
		"ssh_url": "git@github.com:octocat/Hello-World.git",
		"svn_url": "https://svn.github.com/octocat/Hello-World",
		"mirror_url": "git://git.example.com/octocat/Hello-World",
		"homepage": "https://github.com",
		"language": null,
		"forks_count": 9,
		"stargazers_count": 80,
		"watchers_count": 80,
		"size": 108,
		"default_branch": "master",
		"open_issues_count": 0,
		"has_issues": true,
		"has_wiki": true,
		"has_downloads": true,
		"pushed_at": "2011-01-26T19:06:43Z",
		"created_at": "2011-01-26T19:01:12Z",
		"updated_at": "2011-01-26T19:14:43Z",
		"permissions": {
			"admin": false,
			"push": false,
			"pull": true
		}
	}
}`
	object, err := gojson.FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
	} else {
		jsonBytes, _ := gojson.Marshal(object)
		fmt.Println(string(jsonBytes))
		fmt.Println(object.GetDefaultCoding())
	}
}

func Test_json_array_parse111333(t *testing.T) {
	data := `[
  {
    "input_index": 0,
    "candidate_index": 0,
    "delivery_line_1": "1 N Rosedale St",
    "last_line": "Baltimore MD 21229-3737",
    "delivery_point_barcode": "212293737013",
    "components": {
      "primary_number": "1",
      "street_predirection": "N",
      "street_name": "Rosedale",
      "street_suffix": "St",
      "city_name": "Baltimore",
      "state_abbreviation": "MD",
      "zipcode": "21229",
      "plus4_code": "3737",
      "delivery_point": "01",
      "delivery_point_check_digit": "3"
    },
    "metadata": {
      "record_type": "S",
      "zip_type": "Standard",
      "county_fips": "24510",
      "county_name": "Baltimore City",
      "carrier_route": "C047",
      "congressional_district": "07",
      "rdi": "Residential",
      "elot_sequence": "0059",
      "elot_sort": "A",
      "latitude": 39.28602,
      "longitude": -76.6689,
      "precision": "Zip9",
      "time_zone": "Eastern",
      "utc_offset": -5,
      "dst": true
    },
    "analysis": [{
      "dpv_match_code": "Y",
      "dpv_footnotes": "AABB",
      "dpv_cmra": "N",
      "dpv_vacant": "N",
      "active": "Y"
    }]
  },
  {
    "input_index": 0,
    "candidate_index": 1,
    "delivery_line_1": "1 S Rosedale St",
    "last_line": "Baltimore MD 21229-3739",
    "delivery_point_barcode": "212293739011",
    "components": {
      "primary_number": "1",
      "street_predirection": "S",
      "street_name": "Rosedale",
      "street_suffix": "St",
      "city_name": "Baltimore",
      "state_abbreviation": "MD",
      "zipcode": "21229",
      "plus4_code": "3739",
      "delivery_point": "01",
      "delivery_point_check_digit": "1"
    },
    "metadata": {
      "record_type": "S",
      "zip_type": "Standard",
      "county_fips": "24510",
      "county_name": "Baltimore City",
      "carrier_route": "C047",
      "congressional_district": "07",
      "rdi": "Residential",
      "elot_sequence": "0064",
      "elot_sort": "A",
      "latitude": 39.2858,
      "longitude": -76.66889,
      "precision": "Zip9",
      "time_zone": "Eastern",
      "utc_offset": -5,
      "dst": true
    },
    "analysis": [
      {
        "dpv_match_code": "Y",
        "dpv_footnotes": "AABB",
        "dpv_cmra": "N",
        "dpv_vacant": "N",
        "active": "Y"
      }
    ]
  }
]`
	object, err := gojson.FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
	} else {
		jsonBytes, _ := gojson.Marshal(object)
		fmt.Println(string(jsonBytes))
		fmt.Println(object.GetCoding("nihao"))
	}
}