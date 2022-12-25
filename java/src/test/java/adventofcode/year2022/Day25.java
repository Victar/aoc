package adventofcode.year2022;

import adventofcode.BaseTest;
import org.junit.Ignore;
import org.junit.Test;

import java.util.ArrayList;
import java.util.Map;

import static org.junit.Assert.assertEquals;

public class Day25 extends BaseTest {

	public static final int DAY = 25;

	static final Map<Long, Character> digitsToSnafu = Map.of(2L, '2', 1L, '1', 0L, '0', -1L, '-', -2L, '=');

	static final Map<Character, Long> snafuToDigits = Map.of('2', 2L, '1', 1L, '0', 0L, '-', -1L, '=', -2L);

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSingle() throws Exception {
		assertEquals("1", numberToSnafu(1));
		assertEquals("2", numberToSnafu(2));
		assertEquals("1=", numberToSnafu(3));
		assertEquals("1121-1110-1=0", numberToSnafu(314159265));
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		long answer = 0;
		for (final String input : data) {
			answer += snafuToNumber(input);
		}
		System.out.println(numberToSnafu(answer));
	}

	public String numberToSnafu(long number) {
		StringBuilder result = new StringBuilder();
		while (number > 0) {
			long remainder = (number + 2) % 5 - 2;
			result.append(digitsToSnafu.get(remainder));
			number = (number - remainder) / 5;
		}
		return result.reverse().toString();
	}

	public long snafuToNumber(String input) {
		long answer = 0;
		long pow = 1;
		for (int i = 0; i < input.length(); i++) {
			answer += snafuToDigits.get(input.charAt(input.length() - i - 1)) * pow;
			pow = pow * 5;
		}
		return answer;
	}

}
