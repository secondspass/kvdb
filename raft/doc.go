// Package raft is the module that implements all the raft consensus stuff transparently
// and handles the storage on the database and distributing it and getting consensus and
// all that. All interactions with the database must go through raft, which takes care of
// consensus and uses the db module to store the local data
package raft
