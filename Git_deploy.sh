#!/bin/bash

echo $PWD
cd $PWD
git add $PWD/.

# Ask the user for this/her message for commit
echo Enter the message for the commit git :
read message_git
git commit -am "$message_git"
