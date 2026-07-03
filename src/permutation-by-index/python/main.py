import math
import sys

from string import ascii_uppercase, digits

from environs import env

CHARSET: str = env("CHARSET", default=f"{ascii_uppercase}{digits}")
CODE_LENGTH: int = env.int("CODE_LENGTH", default=5)
UNIQUE: bool = env.bool("UNIQUE", default=True)

PERMUTATION: int = env.int("PERMUTATION", default=0)

OUTPUT: str = "Permutation: {code}\n"


def main(
    charset_src: str = CHARSET,
    code_length: int = CODE_LENGTH,
    permutation: int = PERMUTATION,
) -> str:
    if permutation <= 0:
        msg = "Permutation must be greater than zero"
        raise ValueError(msg)

    charset: list[str] = list(charset_src)

    divisors = [1]
    max_permutation = len(charset) - code_length + 1
    for index in range(code_length - 1):
        divisors.append(divisors[index] * (code_length + index))
        max_permutation *= len(charset) - code_length + index + 2

    if permutation > max_permutation:
        msg = "Permutation exceeds the maximum permutation of your charset and code length"
        raise ValueError(msg)

    current_offset: int = permutation
    code_chars: list[str] = []
    for divisor in divisors[::-1]:
        char_index = math.ceil(current_offset / divisor) - 1
        current_offset = (current_offset % divisor) or divisor
        code_chars.append(charset.pop(char_index))

    code = "".join(code_chars)
    sys.stdout.write(OUTPUT.format(code=code))

    return code


if __name__ == "__main__":
    main()
