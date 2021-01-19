# Problem 1 (Use any language)
## The Signal Sounder 4000

A team of engineers have been working hard to trace an issue in a radio system
that has been broadcasting incorrect information. This radio system emits a
series of 32 bit binary numbers very quickly which makes manually looking
for issues painful. We need you to develop a system that allows our clients
to describe what conditions within this array of 32 bit numbers are interesting
to them and point out the locations where these conditions are true.

## The Config

We have a few very important rules that must exist in the system. These rules are
finding where the value in an array is greater than, less than, or equal to a
configurable value. They also need to find where the delta value/index is greater
than, less than, or equal to a configurable value. We also told the client that
a few things they asked for were not feasably possible within the timeframe that
they desired. These features would be finding the beginning or end of a pattern and
composing rules. They agreed to pay 10 times the cost of the software if these rules
were implemented.

The following is an example of the config file format that the client requested.

```json
[
	{
		"id": 1,
		"rule-name": "greater than 10",
		"type": "comparison",
		"check": ">",
		"value": 10
	},
	{
		"id": 2,
		"rule-name": "equal to 10",
		"type": "comparison",
		"check": "=",
		"value": 10
	},
	{
		"id": 3,
		"rule-name": "less than 10",
		"type": "comparison",
		"check": "<",
		"value": 10
	},

	{
		"id": 4,
		"rule-name": "rose more than 5 over 7 numbers",
		"type": "delta",
		"check": ">",
		"change": 5,
		"over": 7
	},
	{
		"id": 5,
		"rule-name": "rose exactly 5 over 7 numbers",
		"type": "delta",
		"check": "=",
		"change": 5,
		"over": 7
	},
	{
		"id": 6,
		"rule-name": "less than 5 over 7 numbers",
		"type": "delta",
		"check": "<",
		"change": 5,
		"over": 7
	},


	// Extra Features they want but don't need

	{
		"id": 7,
		"rule-name": "Rose >= 5 over 7 numbers",
		"type": "composition",
		"all": [
			{
				"id": 8,
				"rule-name": "rose more than 5 over 7 numbers",
				"type": "delta",
				"check": ">",
				"change": 5,
				"over": 7
			},
			{
				"id": 9,
				"rule-name": "rose exactly 5 over 7 numbers",
				"type": "delta",
				"check": "=",
				"change": 5,
				"over": 7
			}
		]
	},
	{
		"id": 10,
		"rule-name": "has [4, 4, 4, 1]",
		"type": "pattern",
		"pattern": [4, 4, 4, 1]
	},

]
```

## The Input

The client has an automated system that takes the radio integer stream and converts
it into a CSV. The following is an example of their data file.

```
tower1
10
12
10
9
1
1
1
1
8
2
3
```

The "index" of these values are as follows:

```
tower1
 0 -> 10
 1 -> 12
 2 -> 10
 3 -> 9
 4 -> 1
 5 -> 1
 6 -> 1
 7 -> 1
 8 -> 8
 9 -> 2
10 -> 3
```

As you can see the "index" increases as you move from the top of the CSV to the bottom.

## The Output

Your program should output in the following format:

```
<rule id>@<triggered index>
10@3
```

Each trigger should be printed on it's own line. All lines should end with `\n` with
`\r\n` being acceptable as well.

## Running 

The program should be runnable in a similar fashion to: 

`./your-submission <config file> <input fule>`. 

We gaurantee that there will be no spaces, special characters, or other
complexities in the file names. No invalid input will be provided. 

# Problem 2 (PHP)

Using modern PHP (think PHP7, composer, autoloading, PSR-4, and best practices) please create a "Routing Class" as defined below.

A routing class is designed to take a url segment and automatically invoke a controller class. It does this by breaking up the url segments that come
in a predictable pattern and invoking the correct class and method and inserting parameters into the method invoked. We will be following
a REST-ful convention and be using common http methods. Your Routing class will be aware of the http method and the actual url segment.

Here is the complete requirement. These tables represents all possible routing logic.

## Unnested routes 

| Method | Segment      | Controller class name | Controller method name | Parameters                                                      |
| ------ | ------------ | --------------------- | ---------------------- | --------------------------------------------------------------- |
| GET    | /patients    | PatientsController    | index                  | none                                                            |
| GET    | /patients/2  | PatientsController    | get                    | this should invoke `get($patientId)` where $patientId = 2       |
| POST   | /patients    | PatientsController    | create                 | none (extra credit for handling the request body)               |
| PATCH  | /patients/2  | PatientsController    | update                 | `update($patientId)`                                            |       
| DELETE | /patients/2  | PatientsController    | delete                 | `delete($patientId)`                                            |


## Nested routes

| Method | Segment                    | Controller class name         | Controller method name | Parameters                                  |
|------- | -------------------------- | ----------------------------- | ---------------------- | ------------------------------------------- |
| GET    | /patients/2/metrics        | PatientsMetricsController     | index                  | `index($patientId)`                         |
| GET    | /patients/2/metrics/abc    | PatientsMetricsController     | get                    | `get($patientId, $metricId)`                |
| POST   | /patients/2/metrics        | PatientsMetricsController     | create                 | `create($patientId)`                        |
| PATCH  | /patients/2/metrics/abc    | PatientsMetricsController     | update                 | `update($patientId, $metricId)`             |       
| DELETE | /patients/2/metrics/abc    | PatientsMetricsController     | delete                 | `delete($patientId, $metricId)`             |


Your Route class should implement the following interface:

```
<?php


namespace App;


/**
 * Interface RouterInterface
 * @package App\Interfaces
 */
interface RouterInterface
{
    /**
     * Declare RESTful resource.
     * @param string $name
     * @return void
     */
    public static function resource($name);

    /**
     * Returns the HTTP request response.
     * @return string
     */
    public function run();
}
```


To invoke the Routing class we'll be calling it statically. For example we'll be defining a `routes.php` that looks like this


```
Route::resource('patients');
Route::resource('patients.metrics');
```

When reviewing your submission, we will be adding more routes and resources to your application. It is expected that the application will be able to handle these without any programatic changes to the routing class


# Problem 3 (Java)

Please navigate [here](https://github.com/healthrecoverysolutions/java-test). 
