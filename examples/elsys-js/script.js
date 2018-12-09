

/** 
 * Decode function is the main entry point in the script,
 * and you should not delete it.
 *
 * Arguments:
 * 'payload' - an array of bytes, raw payload from sensor
 * 'port' - an integer value in range [0, 255], f_port value from sensor
 */

function decode(payload, port) {
    return {
        "integer": 1,
        "string": "string",
        "float": 0.11
    }
}
