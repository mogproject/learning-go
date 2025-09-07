#!/usr/bin/env python3

# imports standard libraries
import os
import argparse

__version__ = '0.0.1'

# paths
SCRIPT_PATH = os.path.realpath(__file__)
SCRIPT_DIR = os.path.dirname(SCRIPT_PATH)
PROJECT_DIR = os.path.dirname(SCRIPT_DIR)


def get_parser():
    """Argument parser."""

    parser = argparse.ArgumentParser(description='{{ program_description }}')
    parser.add_argument('-v', '--version', action='version', version=f'%(prog)s {__version__}')
    parser.add_argument('path', nargs='+')
    return parser


def main(args):
    """Entry point of the program. """

    # main logic
    for p in args.path:
        module_dir = os.path.join(PROJECT_DIR, p)
        if os.path.exists(module_dir):
            print(f'Path already exists: {module_dir}')
            continue

        program_name = os.path.basename(p)

        os.mkdir(module_dir)
        with open(os.path.join(module_dir, f'{program_name}.go'), 'w') as f:
            f.write('\n'.join([
                    'package main\n',
                    'import "fmt"\n',
                    'func main() {\n',
                    '}',
                    ]))
        with open(os.path.join(module_dir, 'go.mod'), 'w') as f:
            f.write('\n'.join([
                    f'module {program_name}\n',
                    'go 1.25.1',
                    ]))
        with open(os.path.join(module_dir, 'Makefile'), 'w') as f:
            f.write('\n'.join([
                '.DEFAULT_GOAL := build',
                'PHONY: fmt vet build\n',
                'fmt:',
                '\tgo fmt ./...\n',
                'vet: fmt',
                '\tgo vet ./...\n',
                'build: vet',
                f'\tgo build -o dist/{program_name}',
                f'\t@echo Created dist/{program_name}',
            ]))
        print(f'Created: {p}')


if __name__ == '__main__':
    main(get_parser().parse_args())
