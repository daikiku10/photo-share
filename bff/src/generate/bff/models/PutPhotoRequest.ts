/* tslint:disable */
/* eslint-disable */
/**
 * BFF
 * photo-share bff
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
 * @interface PutPhotoRequest
 */
export interface PutPhotoRequest {
    /**
     * 投稿写真ID
     * @type {string}
     * @memberof PutPhotoRequest
     */
    id: string;
    /**
     * 写真タイトル
     * @type {string}
     * @memberof PutPhotoRequest
     */
    title: string;
    /**
     * 写真の説明
     * @type {string}
     * @memberof PutPhotoRequest
     */
    description: string;
    /**
     * 写真URL
     * @type {string}
     * @memberof PutPhotoRequest
     */
    imageUrl: string;
    /**
     * 投稿者ID
     * @type {string}
     * @memberof PutPhotoRequest
     */
    authorId: string;
    /**
     * カテゴリID
     * @type {string}
     * @memberof PutPhotoRequest
     */
    categoryId: string;
}

/**
 * Check if a given object implements the PutPhotoRequest interface.
 */
export function instanceOfPutPhotoRequest(value: object): boolean {
    if (!('id' in value)) return false;
    if (!('title' in value)) return false;
    if (!('description' in value)) return false;
    if (!('imageUrl' in value)) return false;
    if (!('authorId' in value)) return false;
    if (!('categoryId' in value)) return false;
    return true;
}

export function PutPhotoRequestFromJSON(json: any): PutPhotoRequest {
    return PutPhotoRequestFromJSONTyped(json, false);
}

export function PutPhotoRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PutPhotoRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'],
        'title': json['title'],
        'description': json['description'],
        'imageUrl': json['imageUrl'],
        'authorId': json['authorId'],
        'categoryId': json['categoryId'],
    };
}

export function PutPhotoRequestToJSON(value?: PutPhotoRequest | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'title': value['title'],
        'description': value['description'],
        'imageUrl': value['imageUrl'],
        'authorId': value['authorId'],
        'categoryId': value['categoryId'],
    };
}

