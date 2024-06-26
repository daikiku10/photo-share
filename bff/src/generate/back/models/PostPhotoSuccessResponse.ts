/* tslint:disable */
/* eslint-disable */
/**
 * API
 * photo-share api
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface PostPhotoSuccessResponse
 */
export interface PostPhotoSuccessResponse {
    /**
     * 登録した写真ID
     * @type {string}
     * @memberof PostPhotoSuccessResponse
     */
    id: string;
}

/**
 * Check if a given object implements the PostPhotoSuccessResponse interface.
 */
export function instanceOfPostPhotoSuccessResponse(value: object): boolean {
    if (!('id' in value)) return false;
    return true;
}

export function PostPhotoSuccessResponseFromJSON(json: any): PostPhotoSuccessResponse {
    return PostPhotoSuccessResponseFromJSONTyped(json, false);
}

export function PostPhotoSuccessResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostPhotoSuccessResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'],
    };
}

export function PostPhotoSuccessResponseToJSON(value?: PostPhotoSuccessResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
    };
}

