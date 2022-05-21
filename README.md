# Neuroscience Data Interface Database

### Introduction
Data Interface Database (DID) is a platform which allows researchers to store, query and exchange metadata and results of analyses in a standardized way. Though the primary goal of this software is to provide a data storage solution for the Neuroscience Data Interface (NDI), it aims to be universal enough so that it can be easily integrated within any data analysis pipelines. The user interacts with the platform using REST API, which is independent of the type of platform or languages. Therefore, the software can be easily integrated with data analysis applications written in any programming languages.

### Data Format
Data is stored in the form of JSON-like document with key-value pairs. It is analogous to an object in Object-oriented programming languages. A document may contain fields that make references to other documents and inherent fields from other documents. It is up to the user to specify the format of their documents - the data types of each of the document's fields and its inherentance relationship with the other documentss. They need to be written in JSON with the required fields. Here is an example of such file.

```json
{
	"classname": "classname",
	"superclasses": [
		{ "name":  "superclass1name" },
		{ "name":  "superclass2name" },
		{ "name":  "superclass3name" }
	],
	"depends_on": [
		{ "name1": "" },
		{ "name2": "" }
	],
	"file": [
		{ "name1": "", "location1": ""},
		{ "name2": "", "location2": ""}
	],
	"field": [
		{
			"name":	"",
			"type":	"",
			"default_value":	"",
			"parameters":		"",
			"queryable":		1,
			"documentation":	""
		},
		{
			"name":	"",
			"type":		"",
			"default_value":	"",
			"parameters":		"",
			"queryable":		1,
			"documentation":	""
		},
		{
			"subfield": {
				"name":		"subfield1",
				"field": [ {
					"name":			"",
					"type":			"",
					"default_value":	"",
					"parameters":		"",
					"queryable":		1,
					"documentation":	""
				} ]
			}
		}
	]
}