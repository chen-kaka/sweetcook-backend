rsync -cavzP --delete-after ./ --exclude-from='.rsync-exclude' root@123.207.1.214:/xy/src/go/src/sweetcook-backend
ssh root@123.207.1.214 "\
  source /etc/profile; \
  cd /xy/src/go/src/sweetcook-backend; \
  go clean -i; \
  go build; \
  sh /xy/src/go/src/sweetcook-backend/run-test-docker.sh \
  "
