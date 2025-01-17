/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
    "/assets": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description Asset list */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["AssetCollection"];
                    };
                };
                500: components["responses"]["ServerError"];
            };
        };
        put?: never;
        /** Create asset */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody: {
                content: {
                    "application/json": components["schemas"]["NewAsset"];
                };
            };
            responses: {
                /** @description Asset created */
                201: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["Asset"];
                    };
                };
                400: components["responses"]["InvalidRequest"];
                404: components["responses"]["NotFound"];
                500: components["responses"]["ServerError"];
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/assets/{assetID}": {
        parameters: {
            query?: never;
            header?: never;
            path: {
                assetID: number;
            };
            cookie?: never;
        };
        get: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    assetID: number;
                };
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description Asset found */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["Asset"];
                    };
                };
                404: components["responses"]["NotFound"];
                500: components["responses"]["ServerError"];
            };
        };
        put: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    assetID: number;
                };
                cookie?: never;
            };
            requestBody: {
                content: {
                    "application/json": components["schemas"]["NewAsset"];
                };
            };
            responses: {
                /** @description Asset updated */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["Asset"];
                    };
                };
                400: components["responses"]["InvalidRequest"];
                404: components["responses"]["NotFound"];
                500: components["responses"]["ServerError"];
            };
        };
        post?: never;
        delete: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    assetID: number;
                };
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description Asset deleted */
                204: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content?: never;
                };
                404: components["responses"]["NotFound"];
                500: components["responses"]["ServerError"];
            };
        };
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/models": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description Model list */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ModelCollection"];
                    };
                };
                500: components["responses"]["ServerError"];
            };
        };
        put?: never;
        /** Create model */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody: {
                content: {
                    "application/json": components["schemas"]["NewModel"];
                };
            };
            responses: {
                /** @description Model created */
                201: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["Model"];
                    };
                };
                400: components["responses"]["InvalidRequest"];
                404: components["responses"]["NotFound"];
                500: components["responses"]["ServerError"];
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/models/{modelID}": {
        parameters: {
            query?: never;
            header?: never;
            path: {
                modelID: number;
            };
            cookie?: never;
        };
        get: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    modelID: number;
                };
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description Model found */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["Model"];
                    };
                };
                404: components["responses"]["NotFound"];
                500: components["responses"]["ServerError"];
            };
        };
        put: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    modelID: number;
                };
                cookie?: never;
            };
            requestBody: {
                content: {
                    "application/json": components["schemas"]["NewModel"];
                };
            };
            responses: {
                /** @description Model updated */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["Model"];
                    };
                };
                400: components["responses"]["InvalidRequest"];
                404: components["responses"]["NotFound"];
                500: components["responses"]["ServerError"];
            };
        };
        post?: never;
        delete: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    modelID: number;
                };
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description Model deleted */
                204: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content?: never;
                };
                400: components["responses"]["InvalidRequest"];
                404: components["responses"]["NotFound"];
                500: components["responses"]["ServerError"];
            };
        };
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/vendors": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description Vendor list */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["VendorCollection"];
                    };
                };
                500: components["responses"]["ServerError"];
            };
        };
        put?: never;
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody: {
                content: {
                    "application/json": components["schemas"]["NewVendor"];
                };
            };
            responses: {
                /** @description Vendor created */
                201: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["Vendor"];
                    };
                };
                400: components["responses"]["InvalidRequest"];
                500: components["responses"]["ServerError"];
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/vendors/{vendorID}": {
        parameters: {
            query?: never;
            header?: never;
            path: {
                vendorID: number;
            };
            cookie?: never;
        };
        get: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    vendorID: number;
                };
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description Vendor found */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["Vendor"];
                    };
                };
                404: components["responses"]["NotFound"];
                500: components["responses"]["ServerError"];
            };
        };
        put?: never;
        post?: never;
        delete: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    vendorID: number;
                };
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description Vendor deleted */
                204: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content?: never;
                };
                400: components["responses"]["InvalidRequest"];
                404: components["responses"]["NotFound"];
                500: components["responses"]["ServerError"];
            };
        };
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/vendors/{vendorID}/models": {
        parameters: {
            query?: never;
            header?: never;
            path: {
                vendorID: number;
            };
            cookie?: never;
        };
        /** List vendor models */
        get: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    vendorID: number;
                };
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description Listed vendor models */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["ModelCollection"];
                    };
                };
                404: components["responses"]["NotFound"];
                500: components["responses"]["ServerError"];
            };
        };
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
}
export type webhooks = Record<string, never>;
export interface components {
    schemas: {
        Asset: {
            /**
             * @description A unique identifier for the asset.
             * @example 12
             */
            id: number;
            /**
             * @description The ID of the model that this asset is an instance of.
             * @example 17
             */
            modelID: number;
            /** @description The asset's serial number. */
            serial?: string;
            /** @description Free text relating to the asset. */
            comments?: string;
            /**
             * Format: date-time
             * @description The instant the vendor was added to the system.
             */
            createdAt: string;
            /**
             * Format: date-time
             * @description The instant the vendor's information was last updated.
             */
            updatedAt: string;
        };
        NewAsset: {
            /**
             * @description The ID of the model that this asset is an instance of.
             * @example 17
             */
            modelID: number;
            /** @description The asset's serial number. */
            serial: string;
            /**
             * @description Free text relating to the asset.
             * @example Only works if facing west.
             */
            comments?: string;
        };
        AssetCollection: {
            items: components["schemas"]["Asset"][];
        };
        Model: {
            /** @description A unique identifier for the model. */
            id: number;
            /** @description The unique vendor-provided identifier for the model. */
            model: string;
            /** @description The ID of the vendor who produces the model. */
            vendorID: number;
            /**
             * @description A readable name for the vendor.
             * @example Acme Inc.
             */
            name: string;
            /**
             * Format: date-time
             * @description The instant the vendor was added to the system.
             */
            createdAt: string;
            /**
             * Format: date-time
             * @description The instant the vendor's information was last updated.
             */
            updatedAt: string;
        };
        NewModel: {
            /** @description The unique vendor-provided identifier for the model. */
            model: string;
            /**
             * @description A readable name for the vendor.
             * @example Acme Inc.
             */
            name?: string;
            /**
             * @description The ID of the vendor that produces the model.
             * @example 36
             */
            vendorID: number;
        };
        Vendor: {
            /** @description A unique identifier for the vendor. */
            id: number;
            /**
             * @description A readable name for the vendor.
             * @example Acme Inc.
             */
            name: string;
            /**
             * Format: date-time
             * @description The instant the vendor was added to the system.
             */
            createdAt: string;
            /**
             * Format: date-time
             * @description The instant the vendor's information was last updated.
             */
            updatedAt: string;
        };
        NewVendor: {
            /**
             * @description A readable name for the vendor.
             * @example Acme Inc.
             */
            name: string;
        };
        ModelCollection: {
            items: components["schemas"]["Model"][];
        };
        VendorCollection: {
            items: components["schemas"]["Vendor"][];
        };
        APIError: {
            /** @description An array of errors associated with specific fields. */
            fields?: components["schemas"]["FieldError"][];
            /**
             * @description A high-level overview of the error condition.
             * @example No widget found with ID 24.
             */
            message?: string;
        };
        FieldError: {
            /**
             * @description The name of a field that failed validation.
             * @example username
             */
            field: string;
            /**
             * @description A description of why the field is invalid.
             * @example This field must be between 1 and 20 characters long.
             */
            message: string;
        };
    };
    responses: {
        /** @description Invalid request */
        InvalidRequest: {
            headers: {
                [name: string]: unknown;
            };
            content: {
                "application/json": components["schemas"]["APIError"];
            };
        };
        /** @description Not found */
        NotFound: {
            headers: {
                [name: string]: unknown;
            };
            content: {
                "application/json": components["schemas"]["APIError"];
            };
        };
        /** @description Server error */
        ServerError: {
            headers: {
                [name: string]: unknown;
            };
            content: {
                "application/json": components["schemas"]["APIError"];
            };
        };
    };
    parameters: never;
    requestBodies: never;
    headers: never;
    pathItems: never;
}
export type $defs = Record<string, never>;
export type operations = Record<string, never>;
