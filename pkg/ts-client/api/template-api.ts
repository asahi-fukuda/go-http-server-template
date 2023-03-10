/* tslint:disable */
/* eslint-disable */
/**
 * Go HTTP Server Template API
 * Go HTTP Server Template API
 *
 * The version of the OpenAPI document: 0.0.1
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import globalAxios, { AxiosPromise, AxiosInstance, AxiosRequestConfig } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from '../common';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { GetOrganization200Response } from '../model';
// @ts-ignore
import { Healthcheck401Response } from '../model';
// @ts-ignore
import { ListOrganizations200Response } from '../model';
/**
 * TemplateApi - axios parameter creator
 * @export
 */
export const TemplateApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * Organization を取得する。
         * @summary Organization 取得
         * @param {string} organizationId Organization ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getOrganization: async (organizationId: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'organizationId' is not null or undefined
            assertParamExists('getOrganization', 'organizationId', organizationId)
            const localVarPath = `/organizations/{organizationId}`
                .replace(`{${"organizationId"}}`, encodeURIComponent(String(organizationId)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * ヘルスチェックを行う。
         * @summary ヘルスチェック
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        healthcheck: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/healthcheck`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * Organization を一覧する。
         * @summary Organization 一覧
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listOrganizations: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/organizations`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * TemplateApi - functional programming interface
 * @export
 */
export const TemplateApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = TemplateApiAxiosParamCreator(configuration)
    return {
        /**
         * Organization を取得する。
         * @summary Organization 取得
         * @param {string} organizationId Organization ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getOrganization(organizationId: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<GetOrganization200Response>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getOrganization(organizationId, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * ヘルスチェックを行う。
         * @summary ヘルスチェック
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async healthcheck(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<void>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.healthcheck(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * Organization を一覧する。
         * @summary Organization 一覧
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async listOrganizations(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ListOrganizations200Response>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.listOrganizations(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * TemplateApi - factory interface
 * @export
 */
export const TemplateApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = TemplateApiFp(configuration)
    return {
        /**
         * Organization を取得する。
         * @summary Organization 取得
         * @param {string} organizationId Organization ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getOrganization(organizationId: string, options?: any): AxiosPromise<GetOrganization200Response> {
            return localVarFp.getOrganization(organizationId, options).then((request) => request(axios, basePath));
        },
        /**
         * ヘルスチェックを行う。
         * @summary ヘルスチェック
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        healthcheck(options?: any): AxiosPromise<void> {
            return localVarFp.healthcheck(options).then((request) => request(axios, basePath));
        },
        /**
         * Organization を一覧する。
         * @summary Organization 一覧
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listOrganizations(options?: any): AxiosPromise<ListOrganizations200Response> {
            return localVarFp.listOrganizations(options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * TemplateApi - object-oriented interface
 * @export
 * @class TemplateApi
 * @extends {BaseAPI}
 */
export class TemplateApi extends BaseAPI {
    /**
     * Organization を取得する。
     * @summary Organization 取得
     * @param {string} organizationId Organization ID
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof TemplateApi
     */
    public getOrganization(organizationId: string, options?: AxiosRequestConfig) {
        return TemplateApiFp(this.configuration).getOrganization(organizationId, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * ヘルスチェックを行う。
     * @summary ヘルスチェック
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof TemplateApi
     */
    public healthcheck(options?: AxiosRequestConfig) {
        return TemplateApiFp(this.configuration).healthcheck(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * Organization を一覧する。
     * @summary Organization 一覧
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof TemplateApi
     */
    public listOrganizations(options?: AxiosRequestConfig) {
        return TemplateApiFp(this.configuration).listOrganizations(options).then((request) => request(this.axios, this.basePath));
    }
}
