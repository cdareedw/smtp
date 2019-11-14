# smtp
Process inbound email and write out headers. 
Used to drop headers to a directory while monitoring for SPF and DKIM compliance. Sophos email appliance allows email to be cloned and sent to another host. We don't want to see the private mail but are interested if the headers are correctly applied.
Key feature was that it didn't bounce or otherwise do anything with the email.


