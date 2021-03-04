import { render, screen } from '@testing-library/react';
import {CreateTree} from './Tree';

test('only root', () => {
    expect(CreateTree("/", [])).toStrictEqual(
        {
            key: "root",
            label: "/",
            nodes: []
        }
    )
})

test('only root with dirs', () => {
    expect(CreateTree("/", ["foo", "bar"])).toStrictEqual(
        {
            key: "root",
            label: "/",
            nodes: [
                {
                    key: "foo",
                    label: "foo"
                },
                {
                    key: "bar",
                    label: "bar"
                },
            ]
        }
    )
})

test('one level', () => {
    expect(CreateTree("/first", [])).toStrictEqual(
        {
            key: "root",
            label: "/",
            nodes: [
                {
                    key: "first",
                    label: "first",
                    nodes: []
                }
            ]
        }
    )
})

test('Two levels', () => {
    expect(CreateTree( "/first/second", [])).toStrictEqual(
        {
            key: "root",
            label: "/",
            nodes: [
                {
                    key: "first",
                    label: "first",
                    nodes: [
                        {
                            key: "second",
                            label: "second",
                            nodes: []
                        }
                    ]
                }
            ]
        }
    )
})

test('Two levels with dirs', () => {
    expect(CreateTree( "/first/second", ["foo", "bar"])).toStrictEqual(
        {
            key: "root",
            label: "/",
            nodes: [
                {
                    key: "first",
                    label: "first",
                    nodes: [
                        {
                            key: "second",
                            label: "second",
                            nodes: [
                                {
                                    key: "foo",
                                    label: "foo",
                                },
                                {
                                    key: "bar",
                                    label: "bar",
                                }

                            ]
                        }
                    ]
                }
            ]
        }
    )
})
