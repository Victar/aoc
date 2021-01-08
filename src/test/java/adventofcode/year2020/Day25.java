package adventofcode.year2020;

import java.util.ArrayList;

import org.junit.Test;

import adventofcode.BaseTest;

import static org.junit.Assert.assertEquals;

public class Day25 extends BaseTest {

	private static final long DIVIDER = 20201227;
	private static final long SUBJECT_NUMBER = 7;

	@Test public void singleCheck() {
		assertEquals(5764801, transformSubjectNumber(8, SUBJECT_NUMBER));
		assertEquals(17807724, transformSubjectNumber(11, SUBJECT_NUMBER));
		assertEquals(11, transformSubjectNumberFindTimes(SUBJECT_NUMBER, 17807724));
		assertEquals(8, transformSubjectNumberFindTimes(SUBJECT_NUMBER, 5764801));
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day25/input.txt");
		long cardExpected = Long.parseLong(data.get(0));//8184785;
		long doorExpected = Long.parseLong(data.get(1));//5293040;
		long cardExpectedTimes = transformSubjectNumberFindTimes(SUBJECT_NUMBER, cardExpected);
		long result = transformSubjectNumber(cardExpectedTimes, doorExpected);
		System.out.println("result:" + result);
	}

	public long transformSubjectNumber(long times, long subjectNumber) {
		long value = 1;
		for (long i = 0; i < times; i++) {
			value = value * subjectNumber;
			value = value % DIVIDER;
		}
		return value;
	}

	public long transformSubjectNumberFindTimes(long subjectNumber, long valueSearch) {
		long value = 1;
		for (long i = 0; true; i++) {
			value = value * subjectNumber;
			value = value % DIVIDER;
			if (value == valueSearch) {
				return ++i;
			}
		}
	}
}
