
# Cash
# Content Addressable Storage for Hypermedia
# (c) 2015-2020 Steve Knoblock
# MIT License

import os
import hashlib

class Cash(object):
    objectStoragePath = "objects" # 'path/to/cas'
    cashSource = "JSON"

    def __init__(self):
        self.yetAnotherProperty = 1
        self.objectStoragePath

    def hello(self, aList):
        print("This is a message inside a class.")
    
    def getObjectId(self, content):
        h = hashlib.sha1()
        h.update(content)
        s = h.hexdigest()
        return s

    def writeObject(self, anObject):

        objectContent = b"cheesy comestibles"

        objectId = self.getObjectId(objectContent)

        #objectId = b"jhsadjfhuaoewwralerjnlafjaosdfhsadsjfhasjdfhlasjdhfajldafhl"

        

        # split a string, mkdir with first part, write file with second part
        shardDirectory = objectId[0:2]
        fileName = objectId[2:38]
        shardPath = os.path.join(self.objectStoragePath, shardDirectory)
        if not os.path.isdir(shardPath):
            os.mkdir(shardPath)
        filePath = os.path.join(shardPath, fileName)
        with open(filePath, "w") as objFile:
            print(objectContent, file=objFile)

obj = Cash()
obj.writeObject(1)
obj.hello([ 1,2,3])

