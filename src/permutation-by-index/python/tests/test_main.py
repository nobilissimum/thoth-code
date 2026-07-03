from unittest import TestCase

import pytest

from main import main


class TestMain(TestCase):
    def test_permutation_first(self) -> None:
        code = main("ABCDE", 3, 1)
        assert code == "ABC"

    def test_permutation_up_multichars(self) -> None:
        code = main("ABCDE", 3, 20)
        assert code == "BDC"

    def test_permutation_last(self) -> None:
        code = main("ABCDE", 3, 60)
        assert code == "EDC"

    def test_permutation_long(self) -> None:
        code = main("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 5, 1)
        assert code == "ABCDE"

    def test_permutation_zero(self) -> None:
        with pytest.raises(ValueError, match="Permutation must be greater than zero"):
            main("ABC", 5, 0)

    def test_permutation_exceeds(self) -> None:
        with pytest.raises(
            ValueError,
            match="Permutation exceeds the maximum permutation of your charset and code length",
        ):
            main("ABC", 5, 61)
