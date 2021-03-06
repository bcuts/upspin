// Copyright 2017 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by upspin gendoc. DO NOT EDIT.
// After editing a command's usage, run 'go generate' to update this file.

/*


The upspin command provides utilities for creating and administering
Upspin files, users, and servers. Although Upspin data is often
accessed through the host file system using upspinfs, the upspin
command is necessary for other tasks, such as: changing a user's
keys (upspin user); updating the wrapped keys after access permissions
are changed (upspin share); or seeing all the information about an
Upspin file beyond what is visible through the host file system
(upspin info). It can also be used separately from upspinfs to
create, read, and update files.

Each subcommand has a -help flag that explains it in more detail.
For instance

	upspin user -help

explains the purpose and usage of the user subcommand.

There is a set of global flags such as -config to identify the
configuration file to use (default $HOME/upspin/config) and -log
to set the logging level for debugging. These flags apply across
the subcommands.

Each subcommand has its own set of flags, which if used must appear
after the subcommand name. For example, to run the ls command with
its -l flag and debugging enabled, run

	upspin -log debug ls -l

For a list of available subcommands and global flags, run

	upspin -help



Usage of upspin:
	upspin [globalflags] <command> [flags] <path>
Upspin commands:
	shell (Interactive mode)
	countersign
	cp
	deletestorage
	deploy
	get
	getref
	info
	keygen
	link
	ls
	mkdir
	put
	repack
	rm
	rotate
	setupdomain
	setupserver
	setupstorage
	setupwriters
	share
	signup
	snapshot
	tar
	user
	whichaccess
Global flags:
  -addr host:port
    	publicly accessible network address (host:port)
  -blocksize size
    	size of blocks when writing large files (default 1048576)
  -cachedir directory
    	directory containing all file caches (default "/Users/r/upspin")
  -cachesize bytes
    	max disk bytes for cache (default 5000000000)
  -config file
    	user's configuration file (default "/Users/r/upspin/config")
  -http address
    	address for incoming insecure network connections (default ":80")
  -https address
    	address for incoming secure network connections (default ":443")
  -kind kind
    	server implementation kind (inprocess, gcp) (default "inprocess")
  -letscache directory
    	Let's Encrypt cache directory
  -log level
    	level of logging: debug, info, error, disabled (default info)
  -project project
    	GCP project name
  -serverconfig value
    	comma-separated list of configuration options (key=value) for this server
  -storeserveruser string
    	user name of the StoreServer
  -tls_cert file
    	TLS Certificate file in PEM format
  -tls_key file
    	TLS Key file in PEM format
  -writethrough
    	make storage cache writethrough



Sub-command countersign

Usage: upspin countersign

Countersign updates the signatures and encrypted data for all items
owned by the user. It is intended to be run after a user has changed
keys.

See the description for rotate for information about updating keys.

Flags:
  -help
    	print more information about the command



Sub-command cp

Usage: upspin cp [opts] file... file or cp [opts] file... directory

Cp copies files into, out of, and within Upspin. If the final
argument is a directory, the files are placed inside it.  The other
arguments must not be directories unless the -R flag is set.

If the final argument is not a directory, cp requires exactly two
path names and copies the contents of the first to the second.
The -R flag requires that the final argument be a directory.

When copying from one Upspin path to another Upspin path, cp can be
very efficient, copying only the references to the data rather than
the data itself.

Flags:
  -R	recursively copy directories
  -help
    	print more information about the command
  -v	log each file as it is copied



Sub-command deletestorage

Usage: upspin deletestorage [-path path... | -ref reference...]

Deletestorage deletes blocks from the store. It is given
either a list of path names, in which case it deletes all blocks
referenced by those names, or a list of references, in which case
it deletes the blocks with those references.

WARNING! Deletestorage is dangerous and should not be used unless
the user can guarantee that the blocks that will be deleted are not
referenced by another path name in any other directory tree, including
snapshots.

Exactly one of the -path or -ref flags must be specified.

For -path, only regular items (not links or directories) can be
processed. Each block will be removed from the store on which it
resides, which in exceptional circumstances may be different from
the user's store.

For -ref, the reference must exactly match the reference's full
value, such as is presented by the info command. The reference is
assumed to refer to the store defined in the user's configuration.

Flags:
  -help
    	print more information about the command
  -path
    	delete all blocks referenced by the path names
  -ref
    	delete individual blocks with the specified references



Sub-command get

Usage: upspin get [-out=outputfile] path

Get writes to standard output the contents identified by the Upspin path.

Flags:
  -help
    	print more information about the command
  -out string
    	output file (default standard output)



Sub-command getref

Usage: upspin getref [-out=outputfile] ref

Getref writes to standard output the contents identified by the reference from
the user's default store server. It does not resolve redirections.

Flags:
  -help
    	print more information about the command
  -out string
    	output file (default standard output)



Sub-command info

Usage: upspin info path...

Info prints to standard output a thorough description of all the
information about named paths, including information provided by
ls but also storage references, sizes, and other metadata.

If the path names an Access or Group file, it is also checked for
validity. If it is a link, the command attempts to access the target
of the link.

Flags:
  -help
    	print more information about the command



Sub-command keygen

Usage: upspin keygen [-curve=256] [-secretseed=seed] [-where=$HOME/.ssh]

Keygen creates a new Upspin key pair and stores the pair in local
files secret.upspinkey and public.upspinkey in $HOME/.ssh. Existing
key pairs are appended to $HOME/.ssh/secret2.upspinkey. Keygen does
not update the information in the key server; use the user -put
command for that.

New users should instead use the signup command to create their
first key. Keygen can be used to create new keys.

See the description for rotate for information about updating keys.

Flags:
  -curve name
    	cryptographic curve name: p256, p384, or p521 (default "p256")
  -help
    	print more information about the command
  -secretseed seed
    	128 bit secret seed in proquint format
  -where directory
    	directory to store keys (default "/Users/r/.ssh")



Sub-command link

Usage: upspin link original_path link_path

Link creates an Upspin link. The link is created at the first path
argument and points to the second path argument.

Flags:
  -help
    	print more information about the command



Sub-command ls

Usage: upspin ls [-l] [path...]

Ls lists the names and, if requested, other properties of Upspin
files and directories. If given no path arguments, it lists the
user's root. By default ls does not follow links; use the -L flag
to learn about the targets of links.

Flags:
  -L	follow links
  -R	recur into subdirectories
  -help
    	print more information about the command
  -l	long format



Sub-command mkdir

Usage: upspin mkdir directory...

Mkdir creates Upspin directories.

Flags:
  -help
    	print more information about the command



Sub-command put

Usage: upspin put [-in=inputfile] path

Put writes its input to the store server and installs a directory
entry with the given path name to refer to the data.

TODO: Delete in favor of cp?

Flags:
  -help
    	print more information about the command
  -in string
    	input file (default standard input)



Sub-command repack

Usage: upspin repack [-pack ee] [flags] path...

Repack rewrites the data referred to by each path , storing it again using the
packing specificied by its -pack option, ee by default. If the data is already
packed with the specified packing, the data is untouched unless the -f (force)
flag is specified, which can be helpful if the data is to be repacked using a
fresh key.

Repack does not delete the old storage. See the deletestorage command
for more information.

Flags:
  -f	force repack even if the file is already packed as requested
  -help
    	print more information about the command
  -pack string
    	packing to use when rewriting (default "ee")
  -r	recur into subdirectories
  -v	verbose: log progress



Sub-command rm

Usage: upspin rm path...

Rm removes Upspin files and directories from the name space.

Rm does not delete the associated storage, which is rarely necessary
or wise: storage can be shared between items and unused storage is
better recovered by automatic means.

See the deletestorage command for more information about deleting
storage.

Flags:
  -help
    	print more information about the command



Sub-command rotate

Usage: upspin rotate

Rotate pushes an updated key to the key server.

To update an Upspin key, the sequence is:

  upspin keygen            # Create new key.
  upspin countersign       # Update file signatures to use new key.
  upspin rotate            # Save new key to key server.
  upspin share -r -fix me@example.com/  # Update keys in file metadata.

Keygen creates a new key and saves the old one. Countersign walks
the file tree and adds signatures with the new key alongside those
for the old. Rotate pushes the new key to the KeyServer. Share walks
the file tree, re-wrapping the encryption keys that were encrypted
with the old key to use the new key.

Some of these steps could be folded together but the full sequence
makes it easier to recover if a step fails.

TODO: Rotate and countersign are terms of art, not clear to users.

Flags:
  -help
    	print more information about the command



Sub-command setupdomain

Usage: upspin setupdomain [-where=$HOME/upspin/deploy] [-cluster] -domain=<name>

Setupdomain is the first step in setting up an upspinserver.
The next steps are 'setupstorage' and 'setupserver'.

It generates keys and config files for Upspin server users, placing them in
$where/$domain (the values of the -where and -domain flags substitute for
$where and $domain respectively) and generates a signature that proves that the
calling Upspin user has control over domain.

If the -cluster flag is specified, keys for upspin-dir@domain and
upspin-store@domain are created instead. This flag should be used when setting
up a domain that will run its directory and store servers separately, requiring
separate users to adminster each one. When -cluster is not specified, keys for
a single user (upspin@domain) are generated.

If any state exists at the given location (-where) then the command aborts.

Flags:
  -cluster
    	generate keys for upspin-dir@domain and upspin-store@domain (default is upspin@domain only)
  -curve name
    	cryptographic curve name: p256, p384, or p521 (default "p256")
  -domain name
    	domain name for this Upspin installation
  -help
    	print more information about the command
  -put-users
    	put server users to the key server
  -where directory
    	directory to store private configuration files (default "/Users/r/upspin/deploy")



Sub-command setupserver

Usage: upspin setupserver -domain=<domain> -host=<host> [-where=$HOME/upspin/deploy] [-writers=user,...]

Setupserver is the final step of setting up an upspinserver.
It assumes that you have run 'setupdomain' and 'setupstorage'.

It registers the user created by 'setupdomain' domain with the key server,
copies the configuration files from $where/$domain to the upspinserver and
restarts it, puts the Writers file, and makes the root for the calling user.

The calling user and the server user are included in the Writers file by
default (giving them write access to the store and directory). You may specify
additional writers with the -writers flag. For instance, if you want all users
@example.com to be able to access storage, specify "-writers=*@example.com".

The calling user must be the same one that ran 'upspin setupdomain'.

Flags:
  -domain name
    	domain name for this Upspin installation
  -help
    	print more information about the command
  -host name
    	host name of upspinserver (empty implies the cluster dir.domain and store.domain)
  -where directory
    	directory to store private configuration files (default "/Users/r/upspin/deploy")
  -writers users
    	additional users to be given write access to this server



Sub-command setupstorage

Usage: upspin -project=<gcp_project_name> setupstorage -domain=<name> <bucket_name>

Setupstorage is the second step in setting up an upspinserver.
The first step is 'setupdomain' and the final step is 'setupserver'.

It creates a Google Cloud Storage bucket and a service account for
accessing that bucket. It then writes the service account private key to
$where/$domain/serviceaccount.json and updates the server
configuration files in that directory to use the specified bucket.

Before running this command, you must create a Google Cloud Project and
associated Billing Account using the Cloud Console:
	https://cloud.google.com/console
The project ID can be any available string, but for clarity it's helpful to
pick something that resembles your domain name.

You must also install the Google Cloud SDK:
	https://cloud.google.com/sdk/downloads
Authenticate and enable the necessary APIs:
	$ gcloud auth login
	$ gcloud --project <project> beta service-management enable iam.googleapis.com storage_api
And, finally, authenticate again in a different way:
	$ gcloud auth application-default login

Flags:
  -domain name
    	domain name for this Upspin installation
  -help
    	print more information about the command
  -where directory
    	directory to store private configuration files (default "/Users/r/upspin/deploy")



Sub-command setupwriters

Usage: upspin setupwriters [-where=$HOME/upspin/deploy] -domain=<domain> <user names>

Setupwriters creates or updates the Writers file for the given domain.
The file lists the names of users granted access to write to the domain's
store server and to create their own root on the directory server.

A wildcard permits access to all users of a domain ("*@example.com").

The user name of the project's directory server is automatically included in
the list, so the directory server can use the store for its own data storage.

Flags:
  -domain name
    	domain name for this Upspin installation
  -help
    	print more information about the command
  -where directory
    	directory containing private configuration files (default "/Users/r/upspin/deploy")



Sub-command share

Usage: upspin share path...

Share reports the user names that have access to each of the argument
paths, and what access rights each has. If the access rights do not
agree with the keys stored in the directory metadata for a path,
that is also reported. Given the -fix flag, share updates the keys
to resolve any such inconsistency. Given both -fix and -force, it
updates the keys regardless. The -d and -r flags apply to directories;
-r states whether the share command should descend into subdirectories.

For the rare case of a world-readable ("read:all") file that is encrypted,
the -unencryptforall flag in combination with -fix will rewrite the file
using the EEIntegrity packing, decrypting it and making its contents
visible to anyone.

See the description for rotate for information about updating keys.

Flags:
  -d	do all files in directory; path must be a directory
  -fix
    	repair incorrect share settings
  -force
    	replace wrapped keys regardless of current state
  -help
    	print more information about the command
  -q	suppress output. Default is to show state for every file
  -r	recur into subdirectories; path must be a directory. assumes -d
  -unencryptforall
    	for currently encrypted read:all files only, rewrite using EEIntegrity; requires -fix or -force



Sub-command signup

Usage: upspin [-config=<file>] signup [flags] <username>

Signup generates an Upspin configuration file and private/public key pair,
stores them locally, and sends a signup request to the public Upspin key server
at key.upspin.io. The server will respond by sending a confirmation email to
the given email address (or "username").

Signup writes a configuration file to $HOME/upspin/config, holding the
username and the location of the directory and store servers. It writes the
public and private keys to $HOME/.ssh. These locations may be set using the
global -config and signup-specific -where flags.

The -dir and -store flags specify the network addresses of the Store and
Directory servers that the Upspin user will use. The -server flag may be used
to specify a single server that acts as both Store and Directory, in which case
the -dir and -store flags must not be set.

By default, signup creates new keys with the p256 cryptographic curve set.
The -curve and -secretseed flags allow the user to control the curve or to
recreate or reuse prior keys.

The -signuponly flag tells signup to skip the generation of the configuration
file and keys and only send the signup request to the key server.

Flags:
  -curve name
    	cryptographic curve name: p256, p384, or p521 (default "p256")
  -dir address
    	Directory server address
  -force
    	create a new user even if keys and config file exist
  -help
    	print more information about the command
  -secretseed seed
    	128 bit secret seed in proquint format
  -server address
    	Store and Directory server address (if combined)
  -signuponly
    	only send signup request to key server; do not generate config or keys
  -store address
    	Store server address
  -where directory
    	directory to store keys (default "/Users/r/.ssh")



Sub-command snapshot

Usage: upspin snapshot

Snapshot requests the system to take a snapshot of the user's
directory tree as soon as possible. Snapshots are created only if
the directory server for the user's root supports them.

Flags:
  -help
    	print more information about the command



Sub-command tar

Usage: upspin tar [-extract [-match prefix -replace substitution] ] upspin_directory local_file

Tar archives an Upspin tree into a local tar file, or with the
-extract flag, unpacks a a local tar file into an Upspin tree.

When extracting, the -match and -replace flags cause the extracted
file to have any prefix that matches be replaced by substitute text.
Whether or not these flags are used, the destination path must
always be in Upspin.

Flags:
  -extract
    	extract from archive
  -help
    	print more information about the command
  -match prefix
    	extract from the archive only those pathnames that match the prefix
  -replace text
    	replace -match prefix with the replacement text
  -v	verbose output



Sub-command user

Usage: upspin user [username...]
              user -put [-in=inputfile] [-force] [username]

User prints in YAML format the user record stored in the key server
for the specified user, by default the current user.

With the -put flag, user writes or replaces the information stored
for the current user, such as to update keys or server information.
The information is read from standard input or from the file provided
with the -in flag. The input must provide the complete record for
the user, and must be in the same YAML format printed by the command
without the -put flag.

When using -put, the command takes no arguments. The name of the
user whose record is to be updated must be provided in the input
record and must either be the current user or the name of another
user whose domain is administered by the current user.

A handy way to use the command is to edit the config file and run
	upspin user | upspin user -put

To install new users see the signup command.

Flags:
  -force
    	force writing user record even if key is empty
  -help
    	print more information about the command
  -in string
    	input file (default standard input)
  -put
    	write new user record



Sub-command whichaccess

Usage: upspin whichaccess path...

Whichaccess reports the Upspin path of the Access file
that controls permissions for each of the argument paths.

Flags:
  -help
    	print more information about the command


*/
package main
