echo off
kubectl.exe --namespace albiondata set image deploy/albiondata-backend albiondata-backend=us.gcr.io/personal-projects-1369/albiondata/backend:%1