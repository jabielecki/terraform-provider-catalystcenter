from flask import Flask, jsonify, request, Response

app = Flask(__name__)


@app.route('/<seg1>/<seg2>/<seg3>/<seg4>/<seg5>', methods=['GET', 'POST'])
def generic_endpoint(seg1, seg2, seg3, seg4, seg5):
    print(f"/{seg1}/{seg2}/{seg3}/{seg4}/{seg5}")
    return Response("Wut is that", 500)


@app.route('/dna/system/api/v1/auth/token', methods=['GET', 'POST'])
def authorization():
    return Response('{"Token": "abcdef"}', status=200)


@app.route('/dna/intent/api/v1/discovery', methods=['GET'])
def get_all_discoveries():
    # Mock data for all discoveries
    response = {
        "response": [
            {
                "id": "10",
                "name": "Discovery Example 1",
                "protocol": "Cisco DNA",
                "ipAddressList": "192.0.2.0, 192.0.2.1"
            },
            {
                "id": "11",
                "name": "Discovery Example 2",
                "protocol": "Cisco DNA",
                "ipAddressList": "192.0.2.2, 192.0.2.3"
            }
        ]
    }
    return jsonify(response)


@app.route('/dna/intent/api/v1/discovery/<discovery_id>', methods=['GET'])
def get_discovery_by_id(discovery_id):
    # Mock data for a specific discovery
    response = {
        "response": {
            "id": discovery_id,
            "name": "Specific Discovery",
            "protocol": "Cisco DNA",
            "ipAddressList": "192.0.2.5"
        }
    }
    return jsonify(response)


@app.route('/dna/intent/api/v1/discovery', methods=['POST'])
def create_discovery():
    # Mock response for creating a discovery
    response = {
        "response": {
            "status": "Discovery created successfully",
            "discoveryId": "12"
        }
    }
    return jsonify(response), 201


@app.route('/dna/intent/api/v1/discovery/<discovery_id>', methods=['DELETE'])
def delete_discovery(discovery_id):
    # Mock response for deleting a discovery
    response = {
        "response": f"Discovery with ID {discovery_id} deleted successfully"
    }
    return jsonify(response), 204


@app.route('/dna/intent/api/v1/discovery/<discovery_id>', methods=['PUT'])
def update_discovery(discovery_id):
    # Mock response for updating a discovery
    response = {
        "response": {
            "status": "Discovery updated successfully",
            "discoveryId": discovery_id
        }
    }
    return jsonify(response)


@app.route('/dna/intent/api/v1/network-device', methods=['GET'])
def get_devices():
    # Mock response for retrieving all network devices
    response = {
        "response": [
            {
                "id": "1",
                "hostname": "Device1",
                "type": "Router",
                "managementIpAddress": "192.0.2.1"
            },
            {
                "id": "2",
                "hostname": "Device2",
                "type": "Switch",
                "managementIpAddress": "192.0.2.2"
            }
        ]
    }
    return jsonify(response)

@app.route('/dna/intent/api/v1/network-device/<device_id>', methods=['GET'])
def get_device_by_id(device_id):
    # Mock response for retrieving a specific network device by ID
    response = {
        "response": {
            "id": device_id,
            "hostname": "Device" + device_id,
            "type": "Router",
            "managementIpAddress": "192.0.2." + device_id
        }
    }
    return jsonify(response)


# Add other endpoints as necessary


if __name__ == '__main__':
    app.run(debug=True)
