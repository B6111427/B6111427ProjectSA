/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntClientEntity,
    EntClientEntityFromJSON,
    EntClientEntityFromJSONTyped,
    EntClientEntityToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntStatusEdges
 */
export interface EntStatusEdges {
    /**
     * Status holds the value of the status edge.
     * @type {Array<EntClientEntity>}
     * @memberof EntStatusEdges
     */
    status?: Array<EntClientEntity>;
}

export function EntStatusEdgesFromJSON(json: any): EntStatusEdges {
    return EntStatusEdgesFromJSONTyped(json, false);
}

export function EntStatusEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStatusEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'status': !exists(json, 'status') ? undefined : ((json['status'] as Array<any>).map(EntClientEntityFromJSON)),
    };
}

export function EntStatusEdgesToJSON(value?: EntStatusEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'status': value.status === undefined ? undefined : ((value.status as Array<any>).map(EntClientEntityToJSON)),
    };
}

