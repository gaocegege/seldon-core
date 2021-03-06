# Model and Deployment Metadata

![metadata](./metadata.svg)


## Incubating feature note
The model metadata feature has currently "incubating" status.
This means that we are currently exploring the best possible interface and functionality for this feature.

As a warning word this means that the API or the way you define metadata may be subject to change before
this feature graduates. If you have any comments or suggestion please open the issue on our GitHub project.


## Examples:
- [Basic Examples for Model with Metadata](../../examples/metadata.html)
- [SKLearn Server example with MinIO](../../examples/minio-sklearn.html)
- [Deployment Level Metadata](../../examples/graph-metadata.html).

- [Deploying models trained with Pachyderm](../../examples/pachyderm.html)
- [Deploying models trained with DVC](../../examples/dvc.html)


## Model Metadata (incubating)

With Seldon you can easily define metadata for your models that is compatible with [kfserving dataplane proposal](https://github.com/kubeflow/kfserving/blob/master/docs/predict-api/v2/required_api.md#model-metadata) specification:
```
$metadata_model_response =
{
  "name" : $string,
  "versions" : [ $string, ... ] #optional,
  "platform" : $string,
  "inputs" : [ $metadata_tensor, ... ],
  "outputs" : [ $metadata_tensor, ... ]
}
```
with
```
$metadata_tensor =
{
  "name" : $string,
  "datatype" : $string,
  "shape" : [ $number, ... ]
}
```

Model metadata can be obtained through GET request at `/api/v1.0/metadata/{MODEL_NAME}` endpoint of your deployment.
Example response:
```json
{
  "name": "my-model",
  "versions": ["my-model/v1"],
  "platform": "platform-name",
  "inputs": [{"datatype": "BYTES", "name": "input", "shape": [1, 5]}],
  "outputs": [{"datatype": "BYTES", "name": "output", "shape": [1, 3]}],
}
```

### Prepackaged model servers

To add metadata to your prepackaged model servers simply add a file named `metadata.yaml`
to the S3 bucket with your model:
```YAML
name: my-model
versions: [my-model/v1]
platform: platform-name
inputs:
- datatype: BYTES
  name: input
  shape: [ 1, 4 ]
outputs:
- datatype: BYTES
  name: output
  shape: [ 3 ]
```

See [SKLearn Server example with MinIO](../../examples/minio-sklearn.html) for more details.


### Python Language Wrapper

You can add model metadata you your custom Python model by implementing `init_metadata` method:

```python
class Model:
    ...
    def init_metadata(self):
        meta = {
            "name": "my-model",
            "versions": ["my-model/v1"],
            "platform": "platform-name",
            "inputs": [{"name": "input", "datatype": "BYTES", "shape": [1, 4]}],
            "outputs": [{"name": "output", "datatype": "BYTES", "shape": [1]}],
        }
        return meta
```

See [Python wrapper](../../python/python_component.html#incubating-features) documentation for more details and
notebook [Basic Examples for Model with Metadata](../../examples/metadata.html).


## Deployment Metadata (incubating)
Model metadata allow you to specify metadata for each of the components (nodes) in your graph.
New orchestrator engine will probe all nodes for their metadata and derive global `inputs` and `outputs` of your graph.
It will then expose them together with all nodes' metadata at a single endpoint `/api/v1.0/metadata/` of your deployment.

![graph-metadata](./graph-metadata.svg)

Example response:
```json
{
  "name": "example",
  "models": {
    "model-1": {
      "name": "Model 1",
      "platform": "platform-name",
      "versions": ["model-version"],
      "inputs": [{"datatype": "BYTES", "name": "input", "shape": [1, 5]}],
      "outputs": [{"datatype": "BYTES", "name": "output", "shape": [1, 3]}]
    },
    "model-2": {
      "name": "Model 2",
      "platform": "platform-name",
      "versions": ["model-version"],
      "inputs": [{"datatype": "BYTES", "name": "input", "shape": [1, 3]}],
      "outputs": [{"datatype": "BYTES", "name": "output", "shape": [3]}]
    }
  },
  "graphinputs": [{"datatype": "BYTES", "name": "input", "shape": [1, 5]}],
  "graphoutputs": [{"datatype": "BYTES", "name": "output", "shape": [3]}]
}
```

See example [notebook](../../examples/graph-metadata.html) for more details.
