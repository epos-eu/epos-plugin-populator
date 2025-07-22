run like

`docker run --network host -v $(pwd)/populate.json:/populate.json epos/epos-plugin-populator populate http://example:33000/api/v1 /populate.json`

with the populate.json file that is something like this:

```json

[
  {
    "version": "main",
    "name": "plugin1",
    "description": "plugin description",
    "version_type": "branch",
    "repository": "https://github.com/somerepository/plugin",
    "runtime": "java",
    "executable": "plugin.jar",
    "arguments": "org.example.com.plugin1",
    "enabled": true,
    "inputFormat": "application/json",
    "outputFormat": "application/epos.geo+json",
    "relations": [
      {
        "relationId": "operation/uid/of/distribution1"
      },
      {
        "relationId": "operation/uid/of/distribution2"
      }
    ]
  },
  {
    "version": "main",
    "name": "plugin2",
    "description": "plugin description",
    "version_type": "branch",
    "repository": "https://github.com/somerepository/plugin2",
    "runtime": "binary",
    "executable": "/build/plugin",
    "arguments": "",
    "enabled": true,
    "inputFormat": "application/json",
    "outputFormat": "application/epos.geo+json",
    "relations": [
      {
        "relationId": "operation/uid/of/distribution1"
      }
    ]
  }
```
