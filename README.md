####ipupdate
######A very simple Dynamic DNS for Cloudflare, written in Go

####How do I use it?

1. Install it! `go get github.com/nickvanw/ipupdate` 

2. Create a configuration file. The example file is fairly clear, you'll need:
	* An API Key and the e-mail address that goes along with it
	* The name of your domain
	* The subdomain you want to update
	* Optional: The numerical ID of the record you wish to update
	
Save it somewhere and use the `-c` to point to it: `iptodate -c /path/to/file.conf`, or save it in `/etc/ipupdate.conf`. 

There are four command line flags you need to know:

* `-c`: points to the path of the config file
* `-u`: updates the record you have configured
* `-a`: If there's no record for the subdomain you've configured, it will make one.
* `-l`: on it's own, it will print out the ID for the subdomain you have configured, paired with `-u`, it allows you to update without having a configured ID.


####Getting Started: 

If you don't already have a record created for the subdomain you wish to update, you can create one by running it with the `-a` flag - be warned, however, that Cloudflare will proxy the record by default, so you'll likely want to update it after you create it, by running with both create and udpate: `-a -u`. This will change the record to point directly at the address. 

Once there is a record in place, you can add the ID to your config file, or use the `-l` flag, which will find the ID from Cloudflare before it updates it. Once you've gotten here, you can put something like:

	ipupdate -l -u -c /home/nick/ipupdate
	
in your cron!

Feel free to fork and submit pull requests or open issues, this was something I made in a couple hours because I wanted something easy to deploy that ran on both Windows and Linux, so it probably has bugs!


