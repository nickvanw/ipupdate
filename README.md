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
	
If you don't already have a record created for the subdomain you wish to update, you can create one by running `ipupdate -a` - be warned, however, that Cloudflare will proxy the record by default (and, AFAIK, there's no way to change that when creating a record), so you'll likely want to update it after you create it, by running `ipupdate -a -u`. This will change the record to point directly at the address. 

Once there is a record in place, you can add the ID to your config file, or use the `-l` flag, which will find the ID from Cloudflare before it updates it.

Feel free to fork and submit pull requests or open issues, this was something I made in a couple hours because I wanted something easy to deploy that ran on both Windows and Linux, so it probably has bugs!


