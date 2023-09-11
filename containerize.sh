#!/bin/bash
#
# Author:   zfoteff
# Version:  v1.0.0
# Create tagged Docker image of Quick Congress project and push to the Docker Repository

_application_version=$(go run main.go --version | egrep -o "([0-9]+.+)$")
_image_registry=https://hub.docker.com/repository/docker/
_image_repository=zfoteff899/quickcongress
_image_tag=$1

echo -e "$(date) [*] Starting Quick-Congress Containerization Process ..."

if [ $# -eq 0 ]
  then
      _image_tag=$(git rev-parse HEAD | cut -c1-6)
fi

VERSION_TAG="${_application_version}-${_image_tag}"
IMAGE_TAG="${_image_repository}:${VERSION_TAG}"

echo -e "$(date) [#] Version: ${VERSION_TAG}"
echo -e "$(date) [#] Image Tag: ${IMAGE_TAG}"

docker login
docker build -t $IMAGE_TAG .
docker push $_image_repository:$VERSION_TAG

echo -e "$(date) [+] Completed Quick-Congress Containerization Process"
