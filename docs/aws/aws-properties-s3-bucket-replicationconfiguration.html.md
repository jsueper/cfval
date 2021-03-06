Amazon S3 ReplicationConfiguration
==================================

`ReplicationConfiguration` is a property of the [AWS::S3::Bucket](aws-properties-s3-bucket.html "AWS::S3::Bucket") resource that specifies replication rules and the AWS Identity and Access Management (IAM) role Amazon Simple Storage Service (Amazon S3) uses to replicate objects.

Syntax
------

``` {.programlisting}
      {
  "Role" : String,
  "Rules" : [ Rule, ... ]
}
    
```

Properties
----------

 `Role`   
The Amazon Resource Name (ARN) of an AWS Identity and Access Management (IAM) role that Amazon S3 assumes when replicating objects. For more information, see [How to Set Up Cross-Region Replication](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr-how-setup.html) in the *Amazon Simple Storage Service Developer Guide*.

*Required*: Yes

*Type*: String

 `Rules`   
A replication rule that specifies which objects to replicate and where they are stored.

*Required*: Yes

*Type*: List of [Amazon S3 ReplicationConfiguration Rules](aws-properties-s3-bucket-replicationconfiguration-rules.html "Amazon S3 ReplicationConfiguration Rules")


