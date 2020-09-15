project_dir=$(cd "$(dirname "$0")" && pwd -P)

openapi_generator=openapitools/openapi-generator-cli:v4.3.1

echo "Project dir: ${project_dir}"

echo "generating client"
rm -rf ${project_dir}/client
mkdir -p ${project_dir}/client
docker run --rm -v ${project_dir}:/local --user $(id -u):$(id -g) ${openapi_generator} generate \
  --git-host "github.com" \
  --git-user-id "crea" \
  --git-repo-id "crea-asia/kala" \
  -i /local/docs/swagger.yaml \
  --model-package kalaclient \
  --api-package kalaclient \
  --invoker-package kalaclient \
  --package-name kalaclient \
  -g go \
  -o /local/client
goimports -w $(find ${project_dir}/client -type f -name '*.go')
rm ${project_dir}/client/go.*
