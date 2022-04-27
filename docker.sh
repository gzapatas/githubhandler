docker_compose="docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -v "$PWD:$PWD" -w="$PWD" docker/compose:1.24.0"
_DEPLOY__COMPOSES="--file docker-compose.yaml"

git checkout $_APP__BRANCH &&
git pull &&
docker image prune -f &&

$docker_compose --project-name $_DEPLOY__PROJECT build --force-rm micros.go.mod &&
$docker_compose --project-name $_DEPLOY__PROJECT $_DEPLOY__COMPOSES build --force-rm $_DEPLOY__SERVICE &&
$docker_compose --project-name $_DEPLOY__PROJECT $_DEPLOY__COMPOSES rm -s -f $_DEPLOY__SERVICE &&
$docker_compose --project-name $_DEPLOY__PROJECT $_DEPLOY__COMPOSES up -d --no-deps $_DEPLOY__SERVICE