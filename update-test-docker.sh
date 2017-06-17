rsync -cavzP --delete-after ./ --exclude-from='.rsync-exclude' root@10.2.124.15:/home/gf/go/src/sweetcook-backend
ssh root@10.2.124.15 "\
  source /etc/profile; \
  cd /home/gf/go/src/sweetcook-backend; \
  go clean -i; \
  go build; \
  sh /home/gf/go/src/sweetcook-backend/run-test-docker.sh \
  "
