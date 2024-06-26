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


import * as runtime from '../runtime';
import type {
  DomainError,
  GetPhotoSuccessResponse,
  PostPhotoRequest,
  PostPhotoSuccessResponse,
  PutPhotoRequest,
  PutPhotoSuccessResponse,
} from '../models/index';
import {
    DomainErrorFromJSON,
    DomainErrorToJSON,
    GetPhotoSuccessResponseFromJSON,
    GetPhotoSuccessResponseToJSON,
    PostPhotoRequestFromJSON,
    PostPhotoRequestToJSON,
    PostPhotoSuccessResponseFromJSON,
    PostPhotoSuccessResponseToJSON,
    PutPhotoRequestFromJSON,
    PutPhotoRequestToJSON,
    PutPhotoSuccessResponseFromJSON,
    PutPhotoSuccessResponseToJSON,
} from '../models/index';

export interface DeletePhotoRequest {
    photoId: string;
}

export interface GetPhotoRequest {
    photoId: string;
}

export interface PostPhotoOperationRequest {
    postPhotoRequest: PostPhotoRequest;
}

export interface PutPhotoOperationRequest {
    photoId: string;
    putPhotoRequest: PutPhotoRequest;
}

/**
 * 
 */
export class PhotoApi extends runtime.BaseAPI {

    /**
     * 投稿写真の1件削除をするAPIです。
     * 投稿写真の1件削除API
     */
    async deletePhotoRaw(requestParameters: DeletePhotoRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['photoId'] == null) {
            throw new runtime.RequiredError(
                'photoId',
                'Required parameter "photoId" was null or undefined when calling deletePhoto().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/photos/{photoId}`.replace(`{${"photoId"}}`, encodeURIComponent(String(requestParameters['photoId']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * 投稿写真の1件削除をするAPIです。
     * 投稿写真の1件削除API
     */
    async deletePhoto(requestParameters: DeletePhotoRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deletePhotoRaw(requestParameters, initOverrides);
    }

    /**
     * 投稿写真情報の1件取得するAPIです。
     * 投稿写真情報の1件取得API
     */
    async getPhotoRaw(requestParameters: GetPhotoRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GetPhotoSuccessResponse>> {
        if (requestParameters['photoId'] == null) {
            throw new runtime.RequiredError(
                'photoId',
                'Required parameter "photoId" was null or undefined when calling getPhoto().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/photos/{photoId}`.replace(`{${"photoId"}}`, encodeURIComponent(String(requestParameters['photoId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GetPhotoSuccessResponseFromJSON(jsonValue));
    }

    /**
     * 投稿写真情報の1件取得するAPIです。
     * 投稿写真情報の1件取得API
     */
    async getPhoto(requestParameters: GetPhotoRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GetPhotoSuccessResponse> {
        const response = await this.getPhotoRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * 投稿写真を登録するAPIです。
     * 投稿写真登録API
     */
    async postPhotoRaw(requestParameters: PostPhotoOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostPhotoSuccessResponse>> {
        if (requestParameters['postPhotoRequest'] == null) {
            throw new runtime.RequiredError(
                'postPhotoRequest',
                'Required parameter "postPhotoRequest" was null or undefined when calling postPhoto().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/photos`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: PostPhotoRequestToJSON(requestParameters['postPhotoRequest']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostPhotoSuccessResponseFromJSON(jsonValue));
    }

    /**
     * 投稿写真を登録するAPIです。
     * 投稿写真登録API
     */
    async postPhoto(requestParameters: PostPhotoOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostPhotoSuccessResponse> {
        const response = await this.postPhotoRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * 投稿写真を編集するAPIです。
     * 投稿写真編集API
     */
    async putPhotoRaw(requestParameters: PutPhotoOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PutPhotoSuccessResponse>> {
        if (requestParameters['photoId'] == null) {
            throw new runtime.RequiredError(
                'photoId',
                'Required parameter "photoId" was null or undefined when calling putPhoto().'
            );
        }

        if (requestParameters['putPhotoRequest'] == null) {
            throw new runtime.RequiredError(
                'putPhotoRequest',
                'Required parameter "putPhotoRequest" was null or undefined when calling putPhoto().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/photos/{photoId}`.replace(`{${"photoId"}}`, encodeURIComponent(String(requestParameters['photoId']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: PutPhotoRequestToJSON(requestParameters['putPhotoRequest']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PutPhotoSuccessResponseFromJSON(jsonValue));
    }

    /**
     * 投稿写真を編集するAPIです。
     * 投稿写真編集API
     */
    async putPhoto(requestParameters: PutPhotoOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PutPhotoSuccessResponse> {
        const response = await this.putPhotoRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
