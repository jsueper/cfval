Fn::FindInMap
=============

The intrinsic function `Fn::FindInMap` returns the value corresponding to keys in a two-level map that is declared in the `Mappings` section.

Declaration
-----------

"Fn::FindInMap" : [ "*`MapName`*", "*`TopLevelKey`*", "*`SecondLevelKey`*"]

Parameters
----------

 MapName   
The logical name of a mapping declared in the Mappings section that contains the keys and values.

 TopLevelKey   
The top-level key name. Its value is a list of key-value pairs.

 SecondLevelKey   
The second-level key name, which is set to one of the keys from the list assigned to *`TopLevelKey`*.

Return Value:
-------------

The value that is assigned to *`SecondLevelKey`*.

Example
-------

The following example shows how to use `Fn::FindInMap` for a template with a `Mappings` section that contains a single map, `RegionMap`, that associates AMIs with AWS regions.

-   The map has 5 top-level keys that correspond to various AWS regions.

-   Each top-level key is assigned a list with two second level keys, `"32"` and `"64"`, that correspond to the AMI's architecture.

-   Each of the second-level keys is assigned an appropriate AMI name.

``` {.programlisting}
      
{
  ...
  "Mappings" : {
    "RegionMap" : {
      "us-east-1" : { "32" : "ami-6411e20d", "64" : "ami-7a11e213" },
      "us-west-1" : { "32" : "ami-c9c7978c", "64" : "ami-cfc7978a" },
      "eu-west-1" : { "32" : "ami-37c2f643", "64" : "ami-31c2f645" },
      "ap-southeast-1" : { "32" : "ami-66f28c34", "64" : "ami-60f28c32" },
      "ap-northeast-1" : { "32" : "ami-9c03a89d", "64" : "ami-a003a8a1" }
    }
  },

  "Resources" : {
     "myEC2Instance" : {
        "Type" : "AWS::EC2::Instance",
        "Properties" : {
           "ImageId" : { "Fn::FindInMap" : [ "RegionMap", { "Ref" : "AWS::Region" }, "32"]},
           "InstanceType" : "m1.small"
        }
     }
 }
}
    
```

The example template contains an `AWS::EC2::Instance` resource whose `ImageId` property is set by the `FindInMap` function.

-   *`MapName`* is set to the map of interest, `"RegionMap"` in this example.

-   *`TopLevelKey`* is set to the region where the stack is created, which is determined by using the `"AWS::Region"` pseudo parameter.

-   *`SecondLevelKey`* is set to the desired architecture, `"32"` for this example.

`FindInMap` returns the AMI assigned to `FindInMap`. For a 32-bit instance in us-east-1, `FindInMap` would return `"ami-6411e20d"`.

Supported Functions
-------------------

You can use the following functions in a `Fn::FindInMap` function:

-   `Fn::FindInMap`

-   `Ref`


