# Agricultural Year Module

## Overview

The Agricultural Year module is the first business capability implemented in
AgroCore.

It represents an agricultural year and establishes the initial architectural
patterns that will guide the implementation of the remaining agricultural
modules.

The module currently supports the creation of agricultural years with optional
opening and closing dates.

## Domain Model

The main domain entity is `AgriculturalYear`.

Current attributes:

| Attribute | Type | Required | Description |
| --- | --- | --- | --- |
| ID | int64 | Persistence generated | Persistence identifier |
| Name | string | Yes | Agricultural year name |
| OpeningDate | Date | No | Opening calendar date |
| ClosingDate | Date | No | Closing calendar date |
| Status | Status | Yes | Current entity status |

A newly created agricultural year starts with:

```text
ID = 0
Status = A